package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "abtest.go",
        "filter.go",
        "preprocess.go",
        "rank.go",
        "recsys.go",
        "relatedrec.go",
        "service.go",
        "store.go",
        "upsrec.go",
    ],
    importpath = "go-common/app/service/bbq/recsys/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/recsys/api/grpc/v1:grpc",
        "//app/service/bbq/recsys/conf",
        "//app/service/bbq/recsys/dao",
        "//app/service/bbq/recsys/model",
        "//app/service/bbq/recsys/service/postprocess",
        "//app/service/bbq/recsys/service/rank",
        "//app/service/bbq/recsys/service/retrieve",
        "//app/service/bbq/recsys/service/util",
        "//library/log",
        "//library/log/infoc",
        "//library/net/trace",
        "//library/stat/prom",
        "//vendor/github.com/go-ego/murmur",
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
        "//app/service/bbq/recsys/service/postprocess:all-srcs",
        "//app/service/bbq/recsys/service/rank:all-srcs",
        "//app/service/bbq/recsys/service/retrieve:all-srcs",
        "//app/service/bbq/recsys/service/util:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
