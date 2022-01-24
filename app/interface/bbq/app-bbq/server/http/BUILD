load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "app.go",
        "comment.go",
        "data.go",
        "feed.go",
        "http.go",
        "location.go",
        "notice.go",
        "push.go",
        "report.go",
        "search.go",
        "share.go",
        "sv.go",
        "topic.go",
        "upload.go",
        "user.go",
    ],
    importpath = "go-common/app/interface/bbq/app-bbq/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/bbq/app-bbq/api/http/v1:http",
        "//app/interface/bbq/app-bbq/conf",
        "//app/interface/bbq/app-bbq/model",
        "//app/interface/bbq/app-bbq/service",
        "//app/interface/bbq/common/auth",
        "//app/interface/bbq/common/http",
        "//app/service/bbq/notice-service/api/v1:api",
        "//app/service/bbq/topic/api",
        "//app/service/bbq/user/api",
        "//app/service/bbq/video/api/grpc/v1:grpc",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/antispam",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/metadata",
        "//library/net/trace",
        "//vendor/github.com/json-iterator/go",
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
