package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "notify_test",
    srcs = [
        "client_test.go",
        "liverpc_client_test.go",
        "notify_test.go",
    ],
    embed = [":notify"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/infra/notify/conf",
        "//app/infra/notify/dao",
        "//app/infra/notify/model",
        "//library/net/rpc/liverpc",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "notify",
    srcs = [
        "client.go",
        "liverpc_client.go",
        "pub.go",
        "sub.go",
    ],
    importpath = "go-common/app/infra/notify/notify",
    tags = ["automanaged"],
    deps = [
        "//app/infra/notify/conf",
        "//app/infra/notify/dao",
        "//app/infra/notify/model",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/netutil",
        "//library/net/rpc/liverpc",
        "//library/stat/prom",
        "//vendor/github.com/Shopify/sarama",
        "//vendor/github.com/bsm/sarama-cluster",
        "//vendor/github.com/rcrowley/go-metrics",
        "//vendor/github.com/snluu/uuid",
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
