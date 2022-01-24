load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "mc_level.go",
    ],
    importpath = "go-common/app/job/live-userexp/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/live-userexp/conf",
        "//app/job/live-userexp/model",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/log",
    ],
)

go_test(
    name = "dao_test",
    srcs = ["mc_level_test.go"],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/live-userexp/conf",
        "//app/job/live-userexp/model",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
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
