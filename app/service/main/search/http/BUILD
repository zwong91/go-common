package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "dm_date.go",
        "dm_search.go",
        "dm_update.go",
        "dmhistory_search.go",
        "http.go",
        "pgc_search.go",
        "pgc_update.go",
        "reply_search.go",
        "reply_update.go",
    ],
    importpath = "go-common/app/service/main/search/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/search/conf",
        "//app/service/main/search/dao",
        "//app/service/main/search/model",
        "//app/service/main/search/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
