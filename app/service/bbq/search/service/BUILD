package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "service.go",
        "video.go",
    ],
    importpath = "go-common/app/service/bbq/search/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/search/api/grpc/v1:grpc",
        "//app/service/bbq/search/conf",
        "//app/service/bbq/search/dao",
        "//app/service/bbq/search/model",
        "//library/ecode",
        "//library/log",
        "//vendor/github.com/json-iterator/go",
        "//vendor/gopkg.in/olivere/elastic.v5:elastic_v5",
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
