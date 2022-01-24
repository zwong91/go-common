package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "elec.go",
        "rank.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/ugcpay-rank/internal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/ugcpay-rank/internal/conf",
        "//app/service/main/ugcpay-rank/internal/dao",
        "//app/service/main/ugcpay-rank/internal/model",
        "//app/service/main/ugcpay-rank/internal/service/rank",
        "//library/log",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/pkg/errors",
        "//vendor/golang.org/x/sync/singleflight",
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
        "//app/service/main/ugcpay-rank/internal/service/rank:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
