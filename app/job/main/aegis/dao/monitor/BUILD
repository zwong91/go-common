package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "monitor",
    srcs = [
        "api.go",
        "dao.go",
        "db.go",
        "redis.go",
    ],
    importpath = "go-common/app/job/main/aegis/dao/monitor",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/aegis/conf",
        "//app/job/main/aegis/model/monitor",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
    name = "monitor_test",
    srcs = [
        "api_test.go",
        "dao_test.go",
        "db_test.go",
        "redis_test.go",
    ],
    embed = [":monitor"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/aegis/conf",
        "//app/job/main/aegis/model/monitor",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)
