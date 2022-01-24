package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "exp",
    srcs = [
        "audit.go",
        "check_param.go",
        "clear_cache.go",
        "level_rule.go",
        "log_debug.go",
        "update_db.go",
        "user_exp.go",
    ],
    importpath = "go-common/app/service/live/xuser/service/exp/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/dao-anchor/api/grpc/v1:grpc",
        "//app/service/live/xuser/api/grpc/v1:grpc",
        "//app/service/live/xuser/conf",
        "//app/service/live/xuser/dao/exp",
        "//app/service/live/xuser/dao/xanchor",
        "//app/service/live/xuser/model/exp",
        "//library/cache",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus/report",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
        "//library/time",
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
