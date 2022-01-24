package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "depend.go",
        "service.go",
        "util.go",
    ],
    importpath = "go-common/app/service/main/dapper-query/service",
    tags = ["manual"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/dapper-query/api/v1:api",
        "//app/service/main/dapper-query/conf",
        "//app/service/main/dapper-query/dao",
        "//app/service/main/dapper-query/model",
        "//app/service/main/dapper-query/pkg/cltclient",
        "//app/service/main/dapper-query/service/opslog",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup",
        "//vendor/github.com/golang/protobuf/ptypes/duration",
        "//vendor/github.com/golang/protobuf/ptypes/timestamp",
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
    srcs = [
        ":package-srcs",
        "//app/service/main/dapper-query/service/opslog:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
