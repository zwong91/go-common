load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "bloomfilter.go",
        "filter.go",
        "parallel.go",
        "ranker.go",
        "recall.go",
        "result.go",
        "score.go",
        "service.go",
        "task.go",
    ],
    importpath = "go-common/app/service/bbq/recsys-recall/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/bbq/recall/proto",
        "//app/service/bbq/recsys-recall/api/grpc/v1:grpc",
        "//app/service/bbq/recsys-recall/conf",
        "//app/service/bbq/recsys-recall/dao",
        "//app/service/bbq/recsys-recall/model",
        "//app/service/bbq/recsys-recall/service/index",
        "//library/log",
        "//vendor/github.com/Dai0522/go-hash/bloomfilter",
        "//vendor/github.com/Dai0522/workpool",
        "//vendor/github.com/golang/protobuf/ptypes/empty",
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
        "//app/service/bbq/recsys-recall/service/index:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
