package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "guard",
    srcs = [
        "clear_cache.go",
        "dahanghai.go",
        "guard.go",
        "log_debug.go",
    ],
    importpath = "go-common/app/service/live/xuser/service/guard/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/room/api/liverpc",
        "//app/service/live/room/api/liverpc/v2:liverpc",
        "//app/service/live/xuser/api/grpc/v1:grpc",
        "//app/service/live/xuser/conf",
        "//app/service/live/xuser/dao/account",
        "//app/service/live/xuser/dao/exp",
        "//app/service/live/xuser/dao/guard",
        "//app/service/live/xuser/model",
        "//app/service/live/xuser/model/dhh",
        "//library/cache",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/queue/databus",
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
