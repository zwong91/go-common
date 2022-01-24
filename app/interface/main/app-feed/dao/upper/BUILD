package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "upper_test",
    srcs = [
        "dao_test.go",
        "redis_test.go",
    ],
    embed = [":upper"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-feed/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "upper",
    srcs = [
        "dao.go",
        "redis.go",
    ],
    importpath = "go-common/app/interface/main/app-feed/dao/upper",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-feed/conf",
        "//app/interface/openplatform/article/model",
        "//app/service/main/archive/api",
        "//app/service/main/feed/model",
        "//app/service/main/feed/rpc/client",
        "//library/cache/redis",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
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
