package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "reply_test",
    srcs = [
        "dao_test.go",
        "reply_test.go",
    ],
    embed = [":reply"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/archive/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "reply",
    srcs = [
        "dao.go",
        "reply.go",
    ],
    importpath = "go-common/app/job/main/archive/dao/reply",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/archive/conf",
        "//library/log",
        "//library/net/http/blademaster",
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
