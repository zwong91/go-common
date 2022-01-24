package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "api.go",
        "cache.go",
        "capsule.go",
        "dao.go",
    ],
    importpath = "go-common/app/job/live/xlottery/internal/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/live/xlottery/internal/conf",
        "//app/job/live/xlottery/internal/model",
        "//app/service/live/rc/api/liverpc",
        "//app/service/live/rc/api/liverpc/v1:liverpc",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/net/rpc/liverpc",
        "//library/queue/databus/report",
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
