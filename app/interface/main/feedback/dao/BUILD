package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "bfs_test.go",
        "dao_test.go",
        "playerCheck_test.go",
        "reply_test.go",
        "session_test.go",
        "tag_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/feedback/conf",
        "//app/interface/main/feedback/model",
        "//library/database/sql",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "bfs.go",
        "dao.go",
        "playerCheck.go",
        "reply.go",
        "session.go",
        "tag.go",
    ],
    importpath = "go-common/app/interface/main/feedback/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/feedback/conf",
        "//app/interface/main/feedback/model",
        "//library/database/sql",
        "//library/log",
        "//library/time",
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
