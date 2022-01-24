package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "backup.go",
        "change-log.go",
        "change-src.go",
        "check-live.go",
        "clear-cache.go",
        "clear-retweet.go",
        "http.go",
        "notify.go",
        "official.go",
        "screen-shot.go",
        "stream.go",
        "stream-cut.go",
        "stream-info.go",
        "stream-mask.go",
        "stream-validate.go",
        "upstream-rtmp.go",
        "upstream-summary.go",
    ],
    importpath = "go-common/app/service/video/stream-mng/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/video/stream-mng/common",
        "//app/service/video/stream-mng/conf",
        "//app/service/video/stream-mng/middleware",
        "//app/service/video/stream-mng/model",
        "//app/service/video/stream-mng/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/metadata",
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
