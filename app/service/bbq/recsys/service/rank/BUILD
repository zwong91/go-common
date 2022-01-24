package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "rank",
    srcs = [
        "feature.go",
        "instance.go",
        "model.go",
        "rank.go",
    ],
    importpath = "go-common/app/service/bbq/recsys/service/rank",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/recsys/api/grpc/v1:grpc",
        "//app/service/bbq/recsys/model",
        "//app/service/bbq/recsys/service/rank/treelite",
        "//app/service/bbq/recsys/service/retrieve",
        "//app/service/bbq/recsys/service/util",
        "//library/log",
        "//vendor/github.com/gogo/protobuf/proto",
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
        "//app/service/bbq/recsys/service/rank/treelite:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
