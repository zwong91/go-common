package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dede_test",
    srcs = ["dao_test.go"],
    embed = [":dede"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//app/service/main/videoup/conf"],
)

go_library(
    name = "dede",
    srcs = [
        "dao.go",
        "redis.go",
    ],
    importpath = "go-common/app/service/main/videoup/dao/dede",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/videoup/conf",
        "//app/service/main/videoup/model/dede",
        "//library/cache/redis",
        "//library/log",
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
