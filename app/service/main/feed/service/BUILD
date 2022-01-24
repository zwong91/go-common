package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "archive_feed_test.go",
        "archive_test.go",
        "bangumi_feed_test.go",
        "feed_test.go",
        "mock_test.go",
        "service_test.go",
        "unread_count_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//app/service/main/feed/conf",
        "//app/service/main/feed/dao",
        "//app/service/main/feed/model",
        "//library/cache/redis",
        "//library/time",
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "archive.go",
        "archive_feed.go",
        "article.go",
        "article_feed.go",
        "bangumi_feed.go",
        "feed.go",
        "new_feed.go",
        "service.go",
        "unread_count.go",
    ],
    importpath = "go-common/app/service/main/feed/service",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/rpc/client",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/feed/conf",
        "//app/service/main/feed/dao",
        "//app/service/main/feed/model",
        "//library/log",
        "//library/sync/errgroup",
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
