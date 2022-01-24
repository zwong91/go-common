package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "collector",
    srcs = [
        "breaker.go",
        "collector.go",
        "detect.go",
        "operation_name.go",
        "process.go",
    ],
    importpath = "go-common/app/service/main/dapper/collector",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/dapper/conf",
        "//app/service/main/dapper/dao",
        "//app/service/main/dapper/model",
        "//app/service/main/dapper/pkg/batchwrite",
        "//app/service/main/dapper/pkg/collect",
        "//app/service/main/dapper/pkg/collect/kafkacollect",
        "//app/service/main/dapper/pkg/collect/tcpcollect",
        "//app/service/main/dapper/pkg/pointwrite",
        "//library/log",
        "//library/net/http/blademaster",
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
    name = "collector_test",
    srcs = [
        "breaker_test.go",
        "collector_test.go",
        "detect_test.go",
        "operation_name_test.go",
    ],
    embed = [":collector"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/dapper/model",
        "//library/net/trace/proto",
    ],
)
