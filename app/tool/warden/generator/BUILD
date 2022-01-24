package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "generator_test",
    srcs = ["genservice_test.go"],
    embed = [":generator"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "generator",
    srcs = [
        "gencscode.go",
        "genservice.go",
        "protoc.go",
    ],
    importpath = "go-common/app/tool/warden/generator",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/warden/generator/templates",
        "//app/tool/warden/types",
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
    srcs = [
        ":package-srcs",
        "//app/tool/warden/generator/templates:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
