// Copyright 2020 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rust

import (
	"github.com/google/blueprint"

	"android/soong/cc"
)

var CovLibraryName = "libprofile-clang-extras"

const profileInstrFlag = "-fprofile-instr-generate=/data/misc/trace/clang-%p-%m.profraw"

type coverage struct {
	Properties cc.CoverageProperties

	// Whether binaries containing this module need --coverage added to their ldflags
	linkCoverage bool
}

func (cov *coverage) props() []interface{} {
	return []interface{}{&cov.Properties}
}

func (cov *coverage) deps(ctx DepsContext, deps Deps) Deps {
	if cov.Properties.NeedCoverageVariant {
		ctx.AddVariationDependencies([]blueprint.Variation{
			{Mutator: "link", Variation: "static"},
		}, cc.CoverageDepTag, CovLibraryName)
	}

	return deps
}

func (cov *coverage) flags(ctx ModuleContext, flags Flags, deps PathDeps) (Flags, PathDeps) {

	if !ctx.DeviceConfig().NativeCoverageEnabled() {
		return flags, deps
	}

	if cov.Properties.CoverageEnabled {
		flags.Coverage = true
		coverage := ctx.GetDirectDepWithTag(CovLibraryName, cc.CoverageDepTag).(cc.LinkableInterface)
		flags.RustFlags = append(flags.RustFlags,
			"-Z instrument-coverage", "-g")
		flags.LinkFlags = append(flags.LinkFlags,
			profileInstrFlag, "-g", coverage.OutputFile().Path().String(), "-Wl,--wrap,open",
			// Upstream LLVM change 6d2d3bd0a6 made
			// -z,start-stop-gc the default.  It drops metadata
			// sections like __llvm_prf_data unless they are marked
			// SHF_GNU_RETAIN.  https://reviews.llvm.org/D97448
			// marks generated sections, including __llvm_prf_data
			// as SHF_GNU_RETAIN.  However this change is not in
			// the Rust toolchain.  Since we link Rust libs with
			// new lld, we should use nostart-stop-gc until the
			// Rust toolchain updates past D97448.
			"-Wl,-z,nostart-stop-gc",
		)
		deps.StaticLibs = append(deps.StaticLibs, coverage.OutputFile().Path())
	}

	return flags, deps
}

func (cov *coverage) begin(ctx BaseModuleContext) {
	if ctx.Host() {
		// Host coverage not yet supported.
	} else {
		// Update useSdk and sdkVersion args if Rust modules become SDK aware.
		cov.Properties = cc.SetCoverageProperties(ctx, cov.Properties, ctx.RustModule().nativeCoverage(), false, "")
	}
}
