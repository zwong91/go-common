load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "testdata",
    srcs = [
        "dao.cache.go",
        "dao.go",
        "multi.go",
        "none.go",
        "single.go",
    ],
    importpath = "go-common/app/tool/cache/testdata",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
    ],
)

go_test(
    name = "testdata_test",
    srcs = [
        "multi_test.go",
        "none_test.go",
        "single_test.go",
    ],
    embed = [":testdata"],
    rundir = ".",
    tags = ["automanaged"],
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
