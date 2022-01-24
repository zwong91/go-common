package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "filter.go",
    ],
    importpath = "go-common/app/interface/bbq/bullet/internal/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/bbq/bullet/api",
        "//app/interface/bbq/bullet/internal/conf",
        "//app/interface/bbq/bullet/internal/model",
        "//app/service/bbq/user/api",
        "//app/service/bbq/video/api/grpc/v1:grpc",
        "//app/service/main/filter/api/grpc/v1:grpc",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/rpc/warden",
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
