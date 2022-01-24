package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "archive.go",
        "play.go",
        "service.go",
        "upload.go",
        "video.go",
    ],
    importpath = "go-common/app/service/bbq/video/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/common/db/bbq",
        "//app/service/bbq/topic/api",
        "//app/service/bbq/video/api/grpc/v1:grpc",
        "//app/service/bbq/video/api/http/v1:http",
        "//app/service/bbq/video/conf",
        "//app/service/bbq/video/dao",
        "//app/service/bbq/video/model",
        "//app/service/bbq/video/model/grpc",
        "//library/conf/env",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup",
        "//vendor/github.com/golang/protobuf/ptypes/empty",
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

go_test(
    name = "service_test",
    srcs = ["service_test.go"],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/service/bbq/video/api/grpc/v1:grpc",
        "//app/service/bbq/video/conf",
    ],
)
