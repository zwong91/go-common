package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "acc.go",
        "cache_delay.go",
        "exp.go",
        "relation.go",
        "service.go",
        "vip.go",
    ],
    importpath = "go-common/app/service/main/account/service",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/account/api",
        "//app/service/main/account/conf",
        "//app/service/main/account/dao",
        "//app/service/main/account/model",
        "//app/service/main/account/model/queue",
        "//app/service/main/coin/api/gorpc",
        "//app/service/main/coin/model",
        "//app/service/main/member/model",
        "//app/service/main/relation/model",
        "//app/service/main/relation/rpc/client",
        "//library/cache/memcache",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/sync/errgroup",
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
    name = "service_test",
    srcs = [
        "acc_test.go",
        "exp_test.go",
        "relation_test.go",
        "service_test.go",
        "vip_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/account/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
