package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "rank",
    srcs = [
        "av.go",
        "av_storage.go",
        "rank.go",
        "tool.go",
        "up.go",
        "up_storage.go",
    ],
    importpath = "go-common/app/service/main/ugcpay-rank/internal/service/rank",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/account/model",
        "//app/service/main/ugcpay-rank/internal/dao",
        "//app/service/main/ugcpay-rank/internal/model",
        "//library/cache/memcache",
        "//library/log",
        "//vendor/github.com/pkg/errors",
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
