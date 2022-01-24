package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "regexp_test",
    srcs = ["regexp_test.go"],
    embed = [":regexp"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/GRbit/go-pcre",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "regexp",
    srcs = [
        "cache.go",
        "regexp.go",
    ],
    importpath = "go-common/app/service/main/filter/service/regexp",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//vendor/github.com/GRbit/go-pcre",
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
