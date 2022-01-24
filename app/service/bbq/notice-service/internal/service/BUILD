package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "push.go",
        "service.go",
    ],
    importpath = "go-common/app/service/bbq/notice-service/internal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/notice-service/api/v1:api",
        "//app/service/bbq/notice-service/internal/conf",
        "//app/service/bbq/notice-service/internal/dao",
        "//app/service/bbq/notice-service/internal/model",
        "//app/service/bbq/push/api/grpc/v1:grpc",
        "//library/log",
        "//library/log/infoc",
        "//library/net/trace",
        "//vendor/github.com/golang/protobuf/ptypes/empty",
        "//vendor/github.com/json-iterator/go",
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
