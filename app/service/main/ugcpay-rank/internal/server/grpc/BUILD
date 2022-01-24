package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "grpc",
    srcs = ["elec.go"],
    importpath = "go-common/app/service/main/ugcpay-rank/internal/server/grpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/ugcpay-rank/api/v1:api",
        "//app/service/main/ugcpay-rank/internal/conf",
        "//app/service/main/ugcpay-rank/internal/model",
        "//app/service/main/ugcpay-rank/internal/service",
        "//app/service/main/ugcpay-rank/internal/service/rank",
        "//library/log",
        "//library/net/rpc/warden",
        "//vendor/github.com/golang/protobuf/ptypes/empty",
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
