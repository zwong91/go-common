package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "opslog",
    srcs = ["opslog.go"],
    importpath = "go-common/app/service/main/dapper-query/service/opslog",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/dapper-query/api/v1:api",
        "//app/service/main/dapper-query/conf",
        "//app/service/main/dapper-query/dao",
        "//app/service/main/dapper-query/model",
        "//app/service/main/dapper-query/pkg/opslog",
        "//app/service/main/dapper-query/util",
        "//library/ecode",
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
