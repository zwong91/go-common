package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "generators_test",
    srcs = ["deepcopy_test.go"],
    embed = [":generators"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//app/tool/gengo/types"],
)

go_library(
    name = "generators",
    srcs = ["deepcopy.go"],
    importpath = "go-common/app/tool/gengo/cmd/deepcopy-gen/generators",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/gengo/args",
        "//app/tool/gengo/cmd/deepcopy-gen/sets",
        "//app/tool/gengo/generator",
        "//app/tool/gengo/namer",
        "//app/tool/gengo/types",
        "//vendor/github.com/golang/glog",
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
