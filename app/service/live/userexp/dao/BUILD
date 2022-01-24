load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "exp.go",
        "exp_log.go",
        "mc_level.go",
    ],
    importpath = "go-common/app/service/live/userexp/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/userexp/conf",
        "//app/service/live/userexp/model",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/log",
        "//library/queue/databus/report",
        "//library/time",
        "//library/xstr",
    ],
)

go_test(
    name = "dao_test",
    srcs = [
        "exp_test.go",
        "mc_level_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/live/userexp/conf",
        "//app/service/live/userexp/model",
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
