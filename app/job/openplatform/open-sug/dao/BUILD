package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "bfs.go",
        "dao.go",
        "image.go",
        "item.go",
        "item_mysql.go",
        "redis.go",
        "season_es.go",
    ],
    importpath = "go-common/app/job/openplatform/open-sug/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/openplatform/sug/model",
        "//app/job/openplatform/open-sug/conf",
        "//app/job/openplatform/open-sug/model",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//vendor/code.google.com/p/graphics-go/graphics",
        "//vendor/github.com/golang/freetype",
        "//vendor/golang.org/x/image/font",
        "//vendor/golang.org/x/image/math/fixed",
        "//vendor/gopkg.in/olivere/elastic.v5:elastic_v5",
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
