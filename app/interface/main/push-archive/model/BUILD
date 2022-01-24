package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "model",
    srcs = [
        "batch_param.go",
        "model.go",
        "xints.go",
    ],
    importpath = "go-common/app/interface/main/push-archive/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/relation/model",
        "//library/log",
        "//library/xstr",
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
    name = "model_test",
    srcs = [
        "batch_param_test.go",
        "xints_test.go",
    ],
    embed = [":model"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/smartystreets/goconvey/convey"],
)
