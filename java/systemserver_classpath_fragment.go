// Copyright 2021 Google Inc. All rights reserved.
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

package java

import (
	"android/soong/android"
	"android/soong/dexpreopt"

	"github.com/google/blueprint"
)

func init() {
	registerSystemserverClasspathBuildComponents(android.InitRegistrationContext)

	android.RegisterSdkMemberType(&systemServerClasspathFragmentMemberType{
		SdkMemberTypeBase: android.SdkMemberTypeBase{
			PropertyName: "systemserverclasspath_fragments",
			SupportsSdk:  true,
		},
	})
}

func registerSystemserverClasspathBuildComponents(ctx android.RegistrationContext) {
	ctx.RegisterModuleType("platform_systemserverclasspath", platformSystemServerClasspathFactory)
	ctx.RegisterModuleType("systemserverclasspath_fragment", systemServerClasspathFactory)
	ctx.RegisterModuleType("prebuilt_systemserverclasspath_fragment", prebuiltSystemServerClasspathModuleFactory)
}

type platformSystemServerClasspathModule struct {
	android.ModuleBase

	ClasspathFragmentBase
}

func platformSystemServerClasspathFactory() android.Module {
	m := &platformSystemServerClasspathModule{}
	initClasspathFragment(m, SYSTEMSERVERCLASSPATH)
	android.InitAndroidArchModule(m, android.DeviceSupported, android.MultilibCommon)
	return m
}

func (p *platformSystemServerClasspathModule) AndroidMkEntries() (entries []android.AndroidMkEntries) {
	return p.classpathFragmentBase().androidMkEntries()
}

func (p *platformSystemServerClasspathModule) GenerateAndroidBuildActions(ctx android.ModuleContext) {
	configuredJars := p.configuredJars(ctx)
	classpathJars := configuredJarListToClasspathJars(ctx, configuredJars, p.classpathType)
	p.classpathFragmentBase().generateClasspathProtoBuildActions(ctx, configuredJars, classpathJars)
}

func (p *platformSystemServerClasspathModule) configuredJars(ctx android.ModuleContext) android.ConfiguredJarList {
	// TODO(satayev): include any apex jars that don't populate their classpath proto config.
	return dexpreopt.GetGlobalConfig(ctx).SystemServerJars
}

type SystemServerClasspathModule struct {
	android.ModuleBase
	android.ApexModuleBase
	android.SdkBase

	ClasspathFragmentBase

	properties systemServerClasspathFragmentProperties

	// Collect the module directory for IDE info in java/jdeps.go.
	modulePaths []string
}

func (s *SystemServerClasspathModule) ShouldSupportSdkVersion(ctx android.BaseModuleContext, sdkVersion android.ApiLevel) error {
	return nil
}

type systemServerClasspathFragmentProperties struct {
	// The contents of this systemserverclasspath_fragment, could be either java_library, or java_sdk_library.
	//
	// The order of this list matters as it is the order that is used in the SYSTEMSERVERCLASSPATH.
	Contents []string
}

func systemServerClasspathFactory() android.Module {
	m := &SystemServerClasspathModule{}
	m.AddProperties(&m.properties)
	android.InitApexModule(m)
	android.InitSdkAwareModule(m)
	initClasspathFragment(m, SYSTEMSERVERCLASSPATH)
	android.InitAndroidArchModule(m, android.DeviceSupported, android.MultilibCommon)
	return m
}

func (s *SystemServerClasspathModule) GenerateAndroidBuildActions(ctx android.ModuleContext) {
	if len(s.properties.Contents) == 0 {
		ctx.PropertyErrorf("contents", "empty contents are not allowed")
	}

	configuredJars := s.configuredJars(ctx)
	classpathJars := configuredJarListToClasspathJars(ctx, configuredJars, s.classpathType)
	s.classpathFragmentBase().generateClasspathProtoBuildActions(ctx, configuredJars, classpathJars)

	// Collect the module directory for IDE info in java/jdeps.go.
	s.modulePaths = append(s.modulePaths, ctx.ModuleDir())
}

func (s *SystemServerClasspathModule) configuredJars(ctx android.ModuleContext) android.ConfiguredJarList {
	global := dexpreopt.GetGlobalConfig(ctx)

	possibleUpdatableModules := gatherPossibleApexModuleNamesAndStems(ctx, s.properties.Contents, systemServerClasspathFragmentContentDepTag)
	jars, unknown := global.ApexSystemServerJars.Filter(possibleUpdatableModules)
	// TODO(satayev): remove geotz ssc_fragment, since geotz is not part of SSCP anymore.
	_, unknown = android.RemoveFromList("geotz", unknown)

	// For non test apexes, make sure that all contents are actually declared in make.
	if global.ApexSystemServerJars.Len() > 0 && len(unknown) > 0 && !android.IsModuleInVersionedSdk(ctx.Module()) {
		ctx.ModuleErrorf("%s in contents must also be declared in PRODUCT_APEX_SYSTEM_SERVER_JARS", unknown)
	}

	return jars
}

type systemServerClasspathFragmentContentDependencyTag struct {
	blueprint.BaseDependencyTag
}

// The systemserverclasspath_fragment contents must never depend on prebuilts.
func (systemServerClasspathFragmentContentDependencyTag) ReplaceSourceWithPrebuilt() bool {
	return false
}

// SdkMemberType causes dependencies added with this tag to be automatically added to the sdk as if
// they were specified using java_systemserver_libs or java_sdk_libs.
func (b systemServerClasspathFragmentContentDependencyTag) SdkMemberType(child android.Module) android.SdkMemberType {
	// If the module is a java_sdk_library then treat it as if it was specified in the java_sdk_libs
	// property, otherwise treat if it was specified in the java_systemserver_libs property.
	if javaSdkLibrarySdkMemberType.IsInstance(child) {
		return javaSdkLibrarySdkMemberType
	}

	return javaSystemserverLibsSdkMemberType
}

func (b systemServerClasspathFragmentContentDependencyTag) ExportMember() bool {
	return true
}

// Contents of system server fragments in an apex are considered to be directly in the apex, as if
// they were listed in java_libs.
func (systemServerClasspathFragmentContentDependencyTag) CopyDirectlyInAnyApex() {}

// Contents of system server fragments require files from prebuilt apex files.
func (systemServerClasspathFragmentContentDependencyTag) RequiresFilesFromPrebuiltApex() {}

var _ android.ReplaceSourceWithPrebuilt = systemServerClasspathFragmentContentDepTag
var _ android.SdkMemberDependencyTag = systemServerClasspathFragmentContentDepTag
var _ android.CopyDirectlyInAnyApexTag = systemServerClasspathFragmentContentDepTag
var _ android.RequiresFilesFromPrebuiltApexTag = systemServerClasspathFragmentContentDepTag

// The tag used for the dependency between the systemserverclasspath_fragment module and its contents.
var systemServerClasspathFragmentContentDepTag = systemServerClasspathFragmentContentDependencyTag{}

func IsSystemServerClasspathFragmentContentDepTag(tag blueprint.DependencyTag) bool {
	return tag == systemServerClasspathFragmentContentDepTag
}

func (s *SystemServerClasspathModule) ComponentDepsMutator(ctx android.BottomUpMutatorContext) {
	module := ctx.Module()
	_, isSourceModule := module.(*SystemServerClasspathModule)

	for _, name := range s.properties.Contents {
		// A systemserverclasspath_fragment must depend only on other source modules, while the
		// prebuilt_systemserverclasspath_fragment_fragment must only depend on other prebuilt modules.
		if !isSourceModule {
			name = android.PrebuiltNameFromSource(name)
		}
		ctx.AddDependency(module, systemServerClasspathFragmentContentDepTag, name)
	}
}

// Collect information for opening IDE project files in java/jdeps.go.
func (s *SystemServerClasspathModule) IDEInfo(dpInfo *android.IdeInfo) {
	dpInfo.Deps = append(dpInfo.Deps, s.properties.Contents...)
	dpInfo.Paths = append(dpInfo.Paths, s.modulePaths...)
}

type systemServerClasspathFragmentMemberType struct {
	android.SdkMemberTypeBase
}

func (s *systemServerClasspathFragmentMemberType) AddDependencies(ctx android.SdkDependencyContext, dependencyTag blueprint.DependencyTag, names []string) {
	ctx.AddVariationDependencies(nil, dependencyTag, names...)
}

func (s *systemServerClasspathFragmentMemberType) IsInstance(module android.Module) bool {
	_, ok := module.(*SystemServerClasspathModule)
	return ok
}

func (s *systemServerClasspathFragmentMemberType) AddPrebuiltModule(ctx android.SdkMemberContext, member android.SdkMember) android.BpModule {
	return ctx.SnapshotBuilder().AddPrebuiltModule(member, "prebuilt_systemserverclasspath_fragment")
}

func (s *systemServerClasspathFragmentMemberType) CreateVariantPropertiesStruct() android.SdkMemberProperties {
	return &systemServerClasspathFragmentSdkMemberProperties{}
}

type systemServerClasspathFragmentSdkMemberProperties struct {
	android.SdkMemberPropertiesBase

	// Contents of the systemserverclasspath fragment
	Contents []string
}

func (s *systemServerClasspathFragmentSdkMemberProperties) PopulateFromVariant(ctx android.SdkMemberContext, variant android.Module) {
	module := variant.(*SystemServerClasspathModule)

	s.Contents = module.properties.Contents
}

func (s *systemServerClasspathFragmentSdkMemberProperties) AddToPropertySet(ctx android.SdkMemberContext, propertySet android.BpPropertySet) {
	builder := ctx.SnapshotBuilder()
	requiredMemberDependency := builder.SdkMemberReferencePropertyTag(true)

	if len(s.Contents) > 0 {
		propertySet.AddPropertyWithTag("contents", s.Contents, requiredMemberDependency)
	}
}

var _ android.SdkMemberType = (*systemServerClasspathFragmentMemberType)(nil)

// A prebuilt version of the systemserverclasspath_fragment module.
type prebuiltSystemServerClasspathModule struct {
	SystemServerClasspathModule
	prebuilt android.Prebuilt
}

func (module *prebuiltSystemServerClasspathModule) Prebuilt() *android.Prebuilt {
	return &module.prebuilt
}

func (module *prebuiltSystemServerClasspathModule) Name() string {
	return module.prebuilt.Name(module.ModuleBase.Name())
}

func (module *prebuiltSystemServerClasspathModule) RequiredFilesFromPrebuiltApex(ctx android.BaseModuleContext) []string {
	return nil
}

var _ android.RequiredFilesFromPrebuiltApex = (*prebuiltSystemServerClasspathModule)(nil)

func prebuiltSystemServerClasspathModuleFactory() android.Module {
	m := &prebuiltSystemServerClasspathModule{}
	m.AddProperties(&m.properties)
	// This doesn't actually have any prebuilt files of its own so pass a placeholder for the srcs
	// array.
	android.InitPrebuiltModule(m, &[]string{"placeholder"})
	android.InitApexModule(m)
	android.InitSdkAwareModule(m)
	android.InitAndroidArchModule(m, android.DeviceSupported, android.MultilibCommon)
	return m
}
