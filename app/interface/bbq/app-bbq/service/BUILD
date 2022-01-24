load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "app.go",
        "comment.go",
        "feed.go",
        "location.go",
        "notice.go",
        "push.go",
        "report.go",
        "search.go",
        "service.go",
        "share.go",
        "sv.go",
        "topic.go",
        "upload.go",
        "user.go",
    ],
    importpath = "go-common/app/interface/bbq/app-bbq/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/bbq/app-bbq/api/http/v1:http",
        "//app/interface/bbq/app-bbq/conf",
        "//app/interface/bbq/app-bbq/dao",
        "//app/interface/bbq/app-bbq/model",
        "//app/interface/bbq/app-bbq/model/grpc",
        "//app/service/bbq/common",
        "//app/service/bbq/notice-service/api/v1:api",
        "//app/service/bbq/recsys/api/grpc/v1:grpc",
        "//app/service/bbq/topic/api",
        "//app/service/bbq/user/api",
        "//app/service/bbq/video/api/grpc/v1:grpc",
        "//library/conf/env",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/rpc/warden",
        "//library/net/trace",
        "//library/time",
        "//vendor/github.com/Dai0522/go-hash/murmur3",
        "//vendor/github.com/golang/protobuf/ptypes/empty",
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
