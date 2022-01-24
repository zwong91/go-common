package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "tcp_test",
    srcs = [
        "conn_test.go",
        "pubsub_test.go",
        "tcp_test.go",
    ],
    embed = [":tcp"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/infra/databus/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "tcp",
    srcs = [
        "conn.go",
        "pubsub.go",
        "tcp.go",
    ],
    importpath = "go-common/app/infra/databus/tcp",
    tags = ["automanaged"],
    deps = [
        "//app/infra/databus/conf",
        "//app/infra/databus/dsn",
        "//app/infra/databus/model",
        "//app/infra/databus/service",
        "//library/conf/env",
        "//library/log",
        "//library/queue/databus",
        "//vendor/github.com/Shopify/sarama",
        "//vendor/github.com/bsm/sarama-cluster",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/rcrowley/go-metrics",
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
