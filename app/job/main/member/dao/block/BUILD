package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
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
        "//app/job/main/member/conf",
        "//app/job/main/member/model/block",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/net/http/blademaster",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "block",
    srcs = [
        "dao.go",
        "http.go",
        "mc.go",
        "mysql.go",
        "notify.go",
    ],
    importpath = "go-common/app/job/main/member/dao/block",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/member/conf",
        "//app/job/main/member/model/block",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
        "//library/net/http/blademaster",
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
