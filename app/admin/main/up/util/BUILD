package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "util",
    srcs = [
        "bits.go",
        "utils.go",
    ],
    importpath = "go-common/app/admin/main/up/util",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/util/timerqueue",
        "//library/log",
        "//library/net/http/blademaster",
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
    srcs = [
        ":package-srcs",
        "//app/admin/main/up/util/databusutil:all-srcs",
        "//app/admin/main/up/util/hbaseutil:all-srcs",
        "//app/admin/main/up/util/mathutil:all-srcs",
        "//app/admin/main/up/util/now:all-srcs",
        "//app/admin/main/up/util/timerqueue:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "util_test",
    srcs = ["utils_test.go"],
    embed = [":util"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/smartystreets/goconvey/convey"],
)
