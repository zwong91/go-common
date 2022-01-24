package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "av.go",
        "bfs.go",
        "dao.go",
        "daoAnchor.go",
        "video.go",
    ],
    importpath = "go-common/app/job/live/dao-anchor-job/internal/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/live/dao-anchor-job/internal/conf",
        "//app/service/live/av/api/liverpc",
        "//app/service/live/av/api/liverpc/v1:liverpc",
        "//app/service/live/dao-anchor/api/grpc/v0:grpc",
        "//app/service/live/dao-anchor/api/grpc/v1:grpc",
        "//app/service/video/stream-mng/api/v1:api",
        "//library/database/bfs",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/rpc/liverpc",
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
