package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "invite_test.go",
        "manager_test.go",
        "medal_mc_test.go",
        "medal_test.go",
        "msg_test.go",
        "pendant_redis_test.go",
        "pendant_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/usersuit/conf",
        "//app/admin/main/usersuit/model",
        "//library/time",
        "//vendor/github.com/satori/go.uuid:go_uuid",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "invite.go",
        "manager.go",
        "medal.go",
        "medal_mc.go",
        "msg.go",
        "pendant.go",
        "pendant_mc.go",
        "pendant_redis.go",
        "upload.go",
    ],
    importpath = "go-common/app/admin/main/usersuit/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/usersuit/conf",
        "//app/admin/main/usersuit/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/sync/errgroup",
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
