package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "esports",
    srcs = [
        "dao.go",
        "message.go",
        "mysql.go",
        "push.go",
    ],
    importpath = "go-common/app/job/main/web-goblin/dao/esports",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/web-goblin/conf",
        "//app/job/main/web-goblin/model/esports",
        "//app/service/main/archive/api",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
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
    name = "esports_test",
    srcs = [
        "dao_test.go",
        "message_test.go",
        "mysql_test.go",
        "push_test.go",
    ],
    embed = [":esports"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/web-goblin/conf",
        "//app/job/main/web-goblin/model/esports",
        "//app/service/main/archive/api",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
