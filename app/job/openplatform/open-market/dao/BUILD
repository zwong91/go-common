package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "berserker.go",
        "dao.go",
        "market_es.go",
        "market_mysql.go",
    ],
    importpath = "go-common/app/job/openplatform/open-market/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/openplatform/open-market/conf",
        "//app/job/openplatform/open-market/model",
        "//library/conf/env",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/pkg/errors",
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
