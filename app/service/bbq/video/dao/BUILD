package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "archive.go",
        "archive_filter.go",
        "bvc.go",
        "check.go",
        "comment.go",
        "dao.cache.go",
        "dao.go",
        "databus.go",
        "video.go",
        "video_bvc.go",
        "video_repository.go",
        "video_upload_process.go",
    ],
    importpath = "go-common/app/service/bbq/video/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/bbq/app-bbq/model",
        "//app/service/bbq/common/db/bbq",
        "//app/service/bbq/video/api/grpc/v1:grpc",
        "//app/service/bbq/video/conf",
        "//app/service/bbq/video/model",
        "//app/service/bbq/video/model/grpc",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//library/cache/redis",
        "//library/conf/env",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/net/rpc/warden",
        "//library/queue/databus",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
        "//library/xstr",
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
