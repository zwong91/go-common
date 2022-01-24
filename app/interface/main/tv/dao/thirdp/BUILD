package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "thirdp_test",
    srcs = [
        "dangbei_test.go",
        "dao_test.go",
        "mango_recom_test.go",
    ],
    embed = [":thirdp"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model/thirdp",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "thirdp",
    srcs = [
        "dangbei.go",
        "dao.go",
        "mango_db.go",
        "mango_mc.go",
        "mango_recom.go",
    ],
    importpath = "go-common/app/interface/main/tv/dao/thirdp",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model/thirdp",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/stat/prom",
        "//library/xstr",
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
