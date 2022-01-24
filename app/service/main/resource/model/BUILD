package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "model_test",
    srcs = ["resource_test.go"],
    embed = [":model"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "model",
    srcs = [
        "abtest.go",
        "ads.go",
        "archive.go",
        "banner.go",
        "const.go",
        "cpm.go",
        "databus.go",
        "relate.go",
        "resource.go",
        "rpc.go",
        "show.go",
    ],
    importpath = "go-common/app/service/main/resource/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-view/model",
        "//app/service/main/archive/api",
        "//library/log",
        "//library/time",
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
