package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "invite_test",
    srcs = [
        "acc_test.go",
        "dao_test.go",
        "mysql_test.go",
        "redis_test.go",
    ],
    embed = [":invite"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/usersuit/conf",
        "//app/service/main/usersuit/model",
        "//vendor/github.com/satori/go.uuid:go_uuid",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "invite",
    srcs = [
        "acc.go",
        "dao.go",
        "mysql.go",
        "redis.go",
    ],
    importpath = "go-common/app/service/main/usersuit/dao/invite",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/usersuit/conf",
        "//app/service/main/usersuit/model",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
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
