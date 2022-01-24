load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = ["user_like_test.go"],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = ["//app/interface/bbq/app-bbq/conf"],
)

go_library(
    name = "dao",
    srcs = [
        "app.go",
        "comment.go",
        "danmu.go",
        "dao.go",
        "feed.go",
        "filter.go",
        "image.go",
        "location.go",
        "notice.go",
        "push.go",
        "redundance.go",
        "report.go",
        "search.go",
        "share.go",
        "sv.go",
        "sv_stat.go",
        "upload.go",
        "user.go",
        "user_like.go",
        "user_relation.go",
        "video.go",
    ],
    importpath = "go-common/app/interface/bbq/app-bbq/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/bbq/app-bbq/api/http/v1:http",
        "//app/interface/bbq/app-bbq/conf",
        "//app/interface/bbq/app-bbq/model",
        "//app/interface/bbq/app-bbq/model/grpc",
        "//app/service/bbq/common",
        "//app/service/bbq/notice-service/api/v1:api",
        "//app/service/bbq/recsys/api/grpc/v1:grpc",
        "//app/service/bbq/user/api",
        "//app/service/bbq/video-image/api/grpc/v1:grpc",
        "//app/service/bbq/video/api/grpc/v1:grpc",
        "//app/service/main/account/api",
        "//app/service/main/filter/api/grpc/v1:grpc",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/net/rpc/warden",
        "//library/sync/pipeline/fanout",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/json-iterator/go",
        "//vendor/google.golang.org/grpc",
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
