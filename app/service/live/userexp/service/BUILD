load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "exp.go",
        "infoc.go",
        "service.go",
    ],
    importpath = "go-common/app/service/live/userexp/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/userexp/conf",
        "//app/service/live/userexp/dao",
        "//app/service/live/userexp/model",
        "//library/cache",
        "//library/log",
    ],
)

go_test(
    name = "service_test",
    srcs = ["exp_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/live/userexp/conf",
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
