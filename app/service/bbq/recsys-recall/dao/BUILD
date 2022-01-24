load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "bloomfilter.go",
        "dao.go",
    ],
    importpath = "go-common/app/service/bbq/recsys-recall/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/recsys-recall/conf",
        "//app/service/bbq/recsys-recall/dao/cache",
        "//library/cache/redis",
        "//vendor/github.com/Dai0522/go-hash/bloomfilter",
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
    srcs = [
        ":package-srcs",
        "//app/service/bbq/recsys-recall/dao/cache:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
