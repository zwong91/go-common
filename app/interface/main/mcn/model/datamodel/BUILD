package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "datamodel",
    srcs = [
        "api.go",
        "tables.go",
    ],
    importpath = "go-common/app/interface/main/mcn/model/datamodel",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/time"],
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
    name = "datamodel_test",
    srcs = ["tables_test.go"],
    embed = [":datamodel"],
    tags = ["automanaged"],
)
