load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "app.go",
        "base.go",
        "comment.go",
        "feed.go",
        "invite.go",
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
    importpath = "go-common/app/interface/bbq/app-bbq/api/http/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/bbq/app-bbq/model",
        "//app/service/bbq/notice-service/api/v1:api",
        "//app/service/bbq/topic/api",
        "//app/service/bbq/user/api",
        "//app/service/bbq/video/api/grpc/v1:grpc",
        "//library/net/http/blademaster",
        "//library/time",
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
