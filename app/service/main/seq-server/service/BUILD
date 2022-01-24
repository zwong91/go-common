package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["service_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/seq-server/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "number.go",
        "seq.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/seq-server/service",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/seq-server/conf",
        "//app/service/main/seq-server/dao",
        "//app/service/main/seq-server/model",
        "//library/ecode",
        "//library/log",
        "//library/net/rpc",
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
