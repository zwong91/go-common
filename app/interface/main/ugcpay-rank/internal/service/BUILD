package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "elec.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/ugcpay-rank/internal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/ugcpay-rank/api/http",
        "//app/interface/main/ugcpay-rank/internal/conf",
        "//app/interface/main/ugcpay-rank/internal/dao",
        "//app/service/main/ugcpay-rank/api/v1:api",
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
