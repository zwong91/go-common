package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "index",
    srcs = [
        "forward_index.go",
        "inverted_index.go",
        "loader.go",
    ],
    importpath = "go-common/app/service/bbq/recsys-recall/service/index",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/bbq/recall/proto",
        "//app/service/bbq/recsys-recall/conf",
        "//library/log",
        "//vendor/github.com/golang/snappy",
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

go_test(
    name = "go_default_test",
    srcs = ["inverted_index_test.go"],
    tags = ["automanaged"],
    deps = ["//app/service/bbq/recsys-recall/service/index:go_default_library"],
)

go_test(
    name = "index_test",
    srcs = ["inverted_index_test.go"],
    deps = [":index"],
)
