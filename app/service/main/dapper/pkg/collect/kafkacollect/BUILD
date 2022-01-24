package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "kafkacollect_test",
    srcs = ["collect_test.go"],
    embed = [":kafkacollect"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/dapper/model",
        "//app/service/main/dapper/pkg/process",
        "//library/log",
    ],
)

go_library(
    name = "kafkacollect",
    srcs = ["collect.go"],
    importpath = "go-common/app/service/main/dapper/pkg/collect/kafkacollect",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/dapper/model",
        "//app/service/main/dapper/pkg/collect",
        "//app/service/main/dapper/pkg/process",
        "//library/log",
        "//library/stat/prom",
        "//vendor/github.com/Shopify/sarama",
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
