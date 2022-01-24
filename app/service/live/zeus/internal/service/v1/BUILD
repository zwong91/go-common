package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = ["zeus.go"],
    importpath = "go-common/app/service/live/zeus/internal/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/common/live/library/mengde",
        "//app/service/live/zeus/api/v1:api",
        "//app/service/live/zeus/expr",
        "//app/service/live/zeus/internal/conf",
        "//app/service/live/zeus/internal/service",
        "//library/ecode",
        "//library/log",
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
