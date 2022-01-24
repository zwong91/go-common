package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "http.go",
        "pub.go",
    ],
    importpath = "go-common/app/infra/databus/http",
    tags = ["automanaged"],
    deps = [
        "//app/infra/databus/conf",
        "//app/infra/databus/dsn",
        "//app/infra/databus/service",
        "//app/infra/databus/tcp",
        "//library/log",
        "//library/net/http/blademaster",
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
