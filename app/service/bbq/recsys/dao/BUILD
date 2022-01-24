package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "bloomfilter.go",
        "dao.go",
        "parallel.go",
        "recall.go",
        "relation.go",
        "user.go",
    ],
    importpath = "go-common/app/service/bbq/recsys/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/recsys-recall/api/grpc/v1:grpc",
        "//app/service/bbq/recsys/api/grpc/v1:grpc",
        "//app/service/bbq/recsys/conf",
        "//app/service/bbq/recsys/dao/parallel",
        "//app/service/bbq/recsys/model",
        "//app/service/bbq/recsys/service/retrieve",
        "//app/service/bbq/search/api/grpc/v1:grpc",
        "//app/service/bbq/user/api",
        "//app/service/main/relation/api",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/net/rpc/warden",
        "//vendor/github.com/Dai0522/go-hash/bloomfilter",
        "//vendor/github.com/Dai0522/workpool",
        "//vendor/github.com/json-iterator/go",
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
        "//app/service/bbq/recsys/dao/parallel:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
