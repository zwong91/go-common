package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "msg",
    srcs = [
        "dao.go",
        "msg.go",
    ],
    importpath = "go-common/app/service/main/videoup/dao/msg",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/videoup/conf",
        "//library/ecode",
        "//library/net/http/blademaster",
        "//library/xstr",
        "//vendor/github.com/pkg/errors",
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
    name = "msg_test",
    srcs = [
        "dao_test.go",
        "msg_test.go",
    ],
    embed = [":msg"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/videoup/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)
