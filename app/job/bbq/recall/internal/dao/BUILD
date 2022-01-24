package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "berserker.go",
        "bloomfilter.go",
        "dao.go",
        "hdfs.go",
        "inverted_index.go",
        "video.go",
    ],
    importpath = "go-common/app/job/bbq/recall/internal/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/bbq/recall/internal/conf",
        "//app/job/bbq/recall/internal/model",
        "//app/job/bbq/recall/proto",
        "//app/job/bbq/recall/proto/quality",
        "//app/service/bbq/recsys-recall/api/grpc/v1:grpc",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/net/rpc/warden",
        "//vendor/github.com/Dai0522/go-hash/bloomfilter",
        "//vendor/github.com/golang/snappy",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
