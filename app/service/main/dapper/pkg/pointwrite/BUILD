package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "pointwrite_test",
    srcs = ["pointwrite_test.go"],
    embed = [":pointwrite"],
    tags = ["automanaged"],
    deps = ["//app/service/main/dapper/model"],
)

go_library(
    name = "pointwrite",
    srcs = ["pointwrite.go"],
    importpath = "go-common/app/service/main/dapper/pkg/pointwrite",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/dapper/model",
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
