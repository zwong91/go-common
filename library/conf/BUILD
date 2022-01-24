load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//library/conf/dsn:all-srcs",
        "//library/conf/env:all-srcs",
        "//library/conf/flagvar:all-srcs",
        "//library/conf/paladin:all-srcs",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "conf",
    srcs = [
        "client.go",
        "client_v2.go",
    ],
    importpath = "go-common/library/conf",
    tags = ["automanaged"],
    deps = [
        "//library/conf/env",
        "//library/ecode",
        "//library/log",
    ],
)

go_test(
    name = "conf_test",
    srcs = [
        "client_test.go",
        "client_v2_test.go",
    ],
    embed = [":conf"],
    rundir = ".",
    tags = ["automanaged"],
)
