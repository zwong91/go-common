package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "email.go",
        "mcn.go",
        "memcache.go",
        "recommend.go",
    ],
    importpath = "go-common/app/job/main/mcn/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/mcn/conf",
        "//app/job/main/mcn/model",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/log",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/pkg/errors",
        "//vendor/gopkg.in/gomail.v2:gomail_v2",
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

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "email_test.go",
        "mcn_test.go",
        "memcache_test.go",
        "recommend_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/mcn/conf",
        "//app/job/main/mcn/model",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
