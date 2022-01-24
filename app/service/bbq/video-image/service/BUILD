package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "score.go",
        "service.go",
        "upload.go",
    ],
    importpath = "go-common/app/service/bbq/video-image/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/video-image/api/grpc/v1:grpc",
        "//app/service/bbq/video-image/api/http/v1:http",
        "//app/service/bbq/video-image/conf",
        "//app/service/bbq/video-image/dao",
        "//library/database/bfs",
        "//library/log",
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
