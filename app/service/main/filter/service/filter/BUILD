package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "filter_test",
    srcs = ["filter_test.go"],
    embed = [":filter"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/filter/conf",
        "//app/service/main/filter/model",
        "//app/service/main/filter/service/area",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "filter",
    srcs = ["filter.go"],
    importpath = "go-common/app/service/main/filter/service/filter",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/filter/conf",
        "//app/service/main/filter/model",
        "//app/service/main/filter/model/actriearea",
        "//app/service/main/filter/service/area",
        "//app/service/main/filter/service/regexp",
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
