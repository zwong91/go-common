package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "gorpc_test",
    srcs = ["archive2_test.go"],
    embed = [":gorpc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
    ],
)

go_library(
    name = "gorpc",
    srcs = [
        "archive1.go",
        "archive2.go",
        "archive3.go",
    ],
    importpath = "go-common/app/service/main/archive/api/gorpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//library/net/rpc",
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
