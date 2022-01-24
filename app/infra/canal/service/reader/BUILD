package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "reader",
    srcs = [
        "offset.go",
        "reader.go",
    ],
    importpath = "go-common/app/infra/canal/service/reader",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//vendor/github.com/Shopify/sarama",
        "//vendor/github.com/pingcap/tidb-tools/tidb_binlog/slave_binlog_proto/go-binlog",
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
