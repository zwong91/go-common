package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "message",
    srcs = [
        "dao.go",
        "send.go",
    ],
    importpath = "go-common/app/admin/main/growup/dao/message",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/growup/conf",
        "//library/log",
        "//library/net/http/blademaster",
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
    name = "message_test",
    srcs = [
        "dao_test.go",
        "send_test.go",
    ],
    embed = [":message"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/growup/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
