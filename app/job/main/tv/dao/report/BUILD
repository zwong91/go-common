package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "report_test",
    srcs = [
        "dao_test.go",
        "mc_test.go",
        "report_test.go",
        "style_test.go",
    ],
    embed = [":report"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/tv/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "report",
    srcs = [
        "dao.go",
        "mc.go",
        "report.go",
        "style.go",
    ],
    importpath = "go-common/app/job/main/tv/dao/report",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/model/pgc",
        "//app/job/main/tv/model/report",
        "//library/cache/memcache",
        "//library/conf/env",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
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
