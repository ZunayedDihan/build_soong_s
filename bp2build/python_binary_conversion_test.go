package bp2build

import (
	"testing"

	"android/soong/android"
	"android/soong/python"
)

func runBp2BuildTestCaseWithLibs(t *testing.T, tc bp2buildTestCase) {
	runBp2BuildTestCase(t, func(ctx android.RegistrationContext) {
		ctx.RegisterModuleType("python_library", python.PythonLibraryFactory)
		ctx.RegisterModuleType("python_library_host", python.PythonLibraryHostFactory)
	}, tc)
}

func TestPythonBinaryHostSimple(t *testing.T) {
	runBp2BuildTestCaseWithLibs(t, bp2buildTestCase{
		description:                        "simple python_binary_host converts to a native py_binary",
		moduleTypeUnderTest:                "python_binary_host",
		moduleTypeUnderTestFactory:         python.PythonBinaryHostFactory,
		moduleTypeUnderTestBp2BuildMutator: python.PythonBinaryBp2Build,
		filesystem: map[string]string{
			"a.py":           "",
			"b/c.py":         "",
			"b/d.py":         "",
			"b/e.py":         "",
			"files/data.txt": "",
		},
		blueprint: `python_binary_host {
    name: "foo",
    main: "a.py",
    srcs: ["**/*.py"],
    exclude_srcs: ["b/e.py"],
    data: ["files/data.txt",],
    libs: ["bar"],
    bazel_module: { bp2build_available: true },
}
    python_library_host {
      name: "bar",
      srcs: ["b/e.py"],
      bazel_module: { bp2build_available: true },
    }`,
		expectedBazelTargets: []string{`py_binary(
    name = "foo",
    data = ["files/data.txt"],
    deps = [":bar"],
    main = "a.py",
    srcs = [
        "a.py",
        "b/c.py",
        "b/d.py",
    ],
)`,
		},
	})
}

func TestPythonBinaryHostPy2(t *testing.T) {
	runBp2BuildTestCaseSimple(t, bp2buildTestCase{
		description:                        "py2 python_binary_host",
		moduleTypeUnderTest:                "python_binary_host",
		moduleTypeUnderTestFactory:         python.PythonBinaryHostFactory,
		moduleTypeUnderTestBp2BuildMutator: python.PythonBinaryBp2Build,
		blueprint: `python_binary_host {
    name: "foo",
    srcs: ["a.py"],
    version: {
        py2: {
            enabled: true,
        },
        py3: {
            enabled: false,
        },
    },

    bazel_module: { bp2build_available: true },
}
`,
		expectedBazelTargets: []string{`py_binary(
    name = "foo",
    python_version = "PY2",
    srcs = ["a.py"],
)`,
		},
	})
}

func TestPythonBinaryHostPy3(t *testing.T) {
	runBp2BuildTestCaseSimple(t, bp2buildTestCase{
		description:                        "py3 python_binary_host",
		moduleTypeUnderTest:                "python_binary_host",
		moduleTypeUnderTestFactory:         python.PythonBinaryHostFactory,
		moduleTypeUnderTestBp2BuildMutator: python.PythonBinaryBp2Build,
		blueprint: `python_binary_host {
    name: "foo",
    srcs: ["a.py"],
    version: {
        py2: {
            enabled: false,
        },
        py3: {
            enabled: true,
        },
    },

    bazel_module: { bp2build_available: true },
}
`,
		expectedBazelTargets: []string{
			// python_version is PY3 by default.
			`py_binary(
    name = "foo",
    srcs = ["a.py"],
)`,
		},
	})
}

func TestPythonBinaryHostArchVariance(t *testing.T) {
	runBp2BuildTestCaseSimple(t, bp2buildTestCase{
		description:                        "test arch variants",
		moduleTypeUnderTest:                "python_binary_host",
		moduleTypeUnderTestFactory:         python.PythonBinaryHostFactory,
		moduleTypeUnderTestBp2BuildMutator: python.PythonBinaryBp2Build,
		filesystem: map[string]string{
			"dir/arm.py": "",
			"dir/x86.py": "",
		},
		blueprint: `python_binary_host {
					 name: "foo-arm",
					 arch: {
						 arm: {
							 srcs: ["arm.py"],
						 },
						 x86: {
							 srcs: ["x86.py"],
						 },
					},
				 }`,
		expectedBazelTargets: []string{
			`py_binary(
    name = "foo-arm",
    srcs = select({
        "//build/bazel/platforms/arch:arm": ["arm.py"],
        "//build/bazel/platforms/arch:x86": ["x86.py"],
        "//conditions:default": [],
    }),
)`,
		},
	})
}
