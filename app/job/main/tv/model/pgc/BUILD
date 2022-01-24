package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "pgc",
    srcs = [
        "content.go",
        "databus.go",
        "license.go",
        "media_cache.go",
        "memcache.go",
        "pgc.go",
        "playurl.go",
        "search.go",
        "style.go",
    ],
    importpath = "go-common/app/job/main/tv/model/pgc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/time"],
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
