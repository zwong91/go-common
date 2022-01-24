package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "util_test",
    srcs = [
        "deviations_test.go",
        "pagination_test.go",
        "string_test.go",
    ],
    embed = [":util"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/stretchr/testify/assert"],
)

go_library(
    name = "util",
    srcs = [
        "deviations.go",
        "pagination.go",
        "rand.go",
        "string.go",
        "time.go",
    ],
    importpath = "go-common/app/service/main/antispam/util",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

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
        "//app/service/main/antispam/util/trie:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
