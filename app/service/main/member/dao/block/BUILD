package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "block",
    srcs = [
        "dao.go",
        "http.go",
        "mc.go",
        "mysql.go",
    ],
    importpath = "go-common/app/service/main/member/dao/block",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/member/conf",
        "//app/service/main/member/model/block",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/pkg/errors",
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
    name = "block_test",
    srcs = [
        "dao_test.go",
        "http_test.go",
        "mc_test.go",
        "mysql_test.go",
    ],
    embed = [":block"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/member/conf",
        "//app/service/main/member/model/block",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/net/http/blademaster",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
