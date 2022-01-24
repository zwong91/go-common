package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "feed_test",
    srcs = [
        "archive_test.go",
        "dao_test.go",
        "rcmd_test.go",
        "tag_test.go",
    ],
    embed = [":feed"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/app/conf",
        "//app/job/main/app/model/feed",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "feed",
    srcs = [
        "archive.go",
        "dao.go",
        "memcache.go",
        "rcmd.go",
        "tag.go",
    ],
    importpath = "go-common/app/job/main/app/dao/feed",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/app/conf",
        "//app/job/main/app/model/feed",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//library/cache/memcache",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
