package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "berserker.go",
        "cmsvideo.go",
        "comment.go",
        "dao.go",
        "email.go",
        "notice.go",
        "user.go",
        "video.go",
    ],
    importpath = "go-common/app/job/bbq/video/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/bbq/video/conf",
        "//app/job/bbq/video/model",
        "//app/service/bbq/notice-service/api/v1:api",
        "//app/service/bbq/search/api/grpc/v1:grpc",
        "//app/service/bbq/sys-msg/api/v1:api",
        "//app/service/bbq/video/api/grpc/v1:grpc",
        "//app/service/main/account/api",
        "//library/cache/redis",
        "//library/conf/env",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/net/rpc/warden",
        "//library/time",
        "//vendor/github.com/json-iterator/go",
        "//vendor/github.com/pkg/errors",
        "//vendor/gopkg.in/gomail.v2:gomail_v2",
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
