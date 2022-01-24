package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "share",
    srcs = [
        "dao.go",
        "mysql.go",
        "redis.go",
    ],
    importpath = "go-common/app/interface/main/web-goblin/dao/share",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/web-goblin/conf",
        "//app/interface/main/web-goblin/model/share",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/stat/prom",
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
