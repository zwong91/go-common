package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "caldiff_test",
    srcs = [
        "cal_diff_test.go",
        "dao_test.go",
        "download_test.go",
        "upbfs_test.go",
    ],
    embed = [":caldiff"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/appstatic/conf",
        "//app/job/main/appstatic/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "caldiff",
    srcs = [
        "cal_diff.go",
        "dao.go",
        "download.go",
        "upbfs.go",
    ],
    importpath = "go-common/app/job/main/appstatic/dao/caldiff",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/appstatic/conf",
        "//app/job/main/appstatic/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
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
