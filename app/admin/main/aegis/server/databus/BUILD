package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "databus",
    srcs = [
        "conf.go",
        "resource.go",
    ],
    importpath = "go-common/app/admin/main/aegis/server/databus",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/aegis/model",
        "//app/admin/main/aegis/model/resource",
        "//library/conf/env",
        "//library/log",
        "//library/queue/databus",
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

go_test(
    name = "databus_test",
    srcs = ["resource_test.go"],
    embed = [":databus"],
    tags = ["automanaged"],
)
