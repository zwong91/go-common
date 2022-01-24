package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dsn_test",
    srcs = [
        "dsn_test.go",
        "example_test.go",
        "query_test.go",
    ],
    embed = [":dsn"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//library/time"],
)

go_library(
    name = "dsn",
    srcs = [
        "doc.go",
        "dsn.go",
        "query.go",
    ],
    importpath = "go-common/library/conf/dsn",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/gopkg.in/go-playground/validator.v9:validator_v9"],
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
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/conf/dsn:go_default_library",
        "//library/time:go_default_library",
    ],
)
