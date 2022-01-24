package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "mc_test.go",
        "mysql_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/identify/conf",
        "//library/database/sql",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "mc.go",
        "mysql.go",
    ],
    importpath = "go-common/app/job/main/identify/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/identify/conf",
        "//app/job/main/identify/model",
        "//app/service/main/passport-auth/model",
        "//library/cache/memcache",
        "//library/database/sql",
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
