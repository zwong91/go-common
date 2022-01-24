package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "actriekey",
    srcs = ["actire.go"],
    importpath = "go-common/app/service/main/filter/model/actriekey",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/log"],
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
