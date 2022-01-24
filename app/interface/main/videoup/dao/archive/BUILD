package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "archive_test",
    srcs = [
        "api_test.go",
        "dao_test.go",
        "redis_test.go",
        "type_test.go",
    ],
    embed = [":archive"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/videoup/conf",
        "//app/interface/main/videoup/model/archive",
        "//app/service/main/up/api/v1:api",
        "//library/ecode",
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "archive",
    srcs = [
        "api.go",
        "dao.go",
        "redis.go",
        "type.go",
    ],
    importpath = "go-common/app/interface/main/videoup/dao/archive",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/dao/tool",
        "//app/interface/main/videoup/conf",
        "//app/interface/main/videoup/model/archive",
        "//app/interface/main/videoup/model/porder",
        "//app/service/main/up/api/v1:api",
        "//library/cache/redis",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/sync/errgroup",
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
