package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "generator",
    srcs = [
        "default_generator.go",
        "default_package.go",
        "doc.go",
        "error_tracker.go",
        "execute.go",
        "generator.go",
        "import_tracker.go",
        "snippet_writer.go",
    ],
    importpath = "go-common/app/tool/gengo/generator",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/gengo/namer",
        "//app/tool/gengo/parser",
        "//app/tool/gengo/types",
        "//vendor/github.com/golang/glog",
        "//vendor/golang.org/x/tools/imports",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["snippet_writer_test.go"],
    tags = ["automanaged"],
    deps = [
        "//app/tool/gengo/generator:go_default_library",
        "//app/tool/gengo/namer:go_default_library",
        "//app/tool/gengo/parser:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "generator_test",
    srcs = ["snippet_writer_test.go"],
    deps = [
        ":generator",
        "//app/tool/gengo/namer",
        "//app/tool/gengo/parser",
    ],
)
