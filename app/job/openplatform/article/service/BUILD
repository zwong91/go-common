package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "retry_test.go",
        "search_test.go",
        "service_test.go",
        "sort_test.go",
        "stat_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/article/model",
        "//app/job/openplatform/article/conf",
        "//app/job/openplatform/article/dao",
        "//library/cache/redis",
        "//library/queue/databus",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "article.go",
        "author.go",
        "cron.go",
        "infoc.go",
        "read.go",
        "recommend.go",
        "search.go",
        "service.go",
        "setting.go",
        "sitemap.go",
        "sort.go",
        "stat.go",
        "tag.go",
    ],
    importpath = "go-common/app/job/openplatform/article/service",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tag/model",
        "//app/interface/main/tag/rpc/client",
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/rpc/client",
        "//app/job/openplatform/article/conf",
        "//app/job/openplatform/article/dao",
        "//app/job/openplatform/article/model",
        "//app/service/main/thumbup/model",
        "//library/conf/env",
        "//library/log",
        "//library/log/infoc",
        "//library/queue/databus",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//vendor/github.com/jaytaylor/html2text",
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
