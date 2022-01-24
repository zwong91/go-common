package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "push_test",
    srcs = [
        "appres_test.go",
        "dao_test.go",
        "push_test.go",
        "redis_test.go",
    ],
    embed = [":push"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-resource/api/v1:api",
        "//app/job/main/appstatic/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "push",
    srcs = [
        "appres.go",
        "dao.go",
        "push.go",
        "redis.go",
    ],
    importpath = "go-common/app/job/main/appstatic/dao/push",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-resource/api/v1:api",
        "//app/job/main/appstatic/conf",
        "//app/job/main/appstatic/model",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/naming/discovery",
        "//library/net/http/blademaster",
        "//library/net/rpc/warden",
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
