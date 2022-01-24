load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "exp.go",
        "service.go",
    ],
    importpath = "go-common/app/job/live-userexp/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/live-userexp/conf",
        "//app/job/live-userexp/dao",
        "//app/job/live-userexp/model",
        "//library/log",
        "//library/queue/databus",
    ],
)

go_test(
    name = "service_test",
    srcs = ["exp_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/live-userexp/conf",
        "//app/job/live-userexp/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
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
