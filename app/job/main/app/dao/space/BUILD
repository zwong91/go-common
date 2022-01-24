package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "space_test",
    srcs = [
        "api_test.go",
        "dao_test.go",
        "redis_test.go",
    ],
    embed = [":space"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/app/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "space",
    srcs = [
        "api.go",
        "dao.go",
        "redis.go",
    ],
    importpath = "go-common/app/job/main/app/dao/space",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/rpc/client",
        "//app/job/main/app/conf",
        "//app/job/main/app/model",
        "//app/job/main/app/model/space",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//library/cache/redis",
        "//library/ecode",
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
