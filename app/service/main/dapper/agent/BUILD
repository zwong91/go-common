package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "agent",
    srcs = ["agent.go"],
    importpath = "go-common/app/service/main/dapper/agent",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/dapper/conf",
        "//app/service/main/dapper/pkg/deliver",
        "//app/service/main/dapper/pkg/diskqueue",
        "//app/service/main/dapper/server/udpcollect",
        "//library/log",
        "//library/net/trace/proto",
        "//vendor/github.com/golang/protobuf/proto",
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
