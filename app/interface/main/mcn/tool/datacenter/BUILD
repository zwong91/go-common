package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "datacenter_test",
    srcs = [
        "http_client_test.go",
        "sqltool_test.go",
    ],
    embed = [":datacenter"],
    tags = ["automanaged"],
    deps = ["//library/log"],
)

go_library(
    name = "datacenter",
    srcs = [
        "http_client.go",
        "sqltool.go",
    ],
    importpath = "go-common/app/interface/main/mcn/tool/datacenter",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//library/stat",
        "//vendor/github.com/pkg/errors",
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
