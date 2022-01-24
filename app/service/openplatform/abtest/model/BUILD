package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "model_test",
    srcs = ["stra_test.go"],
    embed = [":model"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/smartystreets/goconvey/convey"],
)

go_library(
    name = "model",
    srcs = [
        "abtest.go",
        "group.go",
        "stra.go",
        "version.go",
    ],
    importpath = "go-common/app/service/openplatform/abtest/model",
    tags = ["automanaged"],
    deps = [
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
    srcs = [
        ":package-srcs",
        "//app/service/openplatform/abtest/model/jump:all-srcs",
        "//app/service/openplatform/abtest/model/validator:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
