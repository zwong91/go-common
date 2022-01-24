package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "region",
    srcs = ["region.go"],
    importpath = "go-common/app/interface/main/app-tag/model/region",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-tag/model",
        "//app/interface/main/app-tag/model/bangumi",
        "//app/interface/main/app-tag/model/tag",
        "//app/service/main/archive/api",
        "//library/time",
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
