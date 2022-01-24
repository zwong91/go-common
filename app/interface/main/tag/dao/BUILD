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
        "grpc_channel_test.go",
        "grpc_tag_test.go",
        "rpc_test.go",
        "spam_redis_test.go",
        "tagarc_prid_redis_test.go",
        "tagarc_redis_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tag/conf",
        "//app/service/main/archive/api",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "grpc_channel.go",
        "grpc_tag.go",
        "rpc.go",
        "spam_redis.go",
        "tagarc_prid_redis.go",
        "tagarc_redis.go",
    ],
    importpath = "go-common/app/interface/main/tag/dao",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tag/conf",
        "//app/interface/main/tag/model",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//app/service/main/tag/api",
        "//library/cache/redis",
        "//library/log",
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
