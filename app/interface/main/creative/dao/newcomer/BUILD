package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "newcomer_test",
    srcs = [
        "api_test.go",
        "dao_test.go",
        "redis_test.go",
        "task_test.go",
    ],
    embed = [":newcomer"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/newcomer",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "newcomer",
    srcs = [
        "api.go",
        "dao.go",
        "redis.go",
        "task.go",
    ],
    importpath = "go-common/app/interface/main/creative/dao/newcomer",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/tool",
        "//app/interface/main/creative/model/newcomer",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
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
