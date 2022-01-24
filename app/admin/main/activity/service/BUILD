package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "archive.go",
        "likes.go",
        "service.go",
        "subject.go",
    ],
    importpath = "go-common/app/admin/main/activity/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/activity/conf",
        "//app/admin/main/activity/dao",
        "//app/admin/main/activity/model",
        "//app/interface/main/tag/model",
        "//app/interface/main/tag/rpc/client",
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/rpc/client",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/jinzhu/gorm",
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
    srcs = [
        ":package-srcs",
        "//app/admin/main/activity/service/kfc:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "service_test",
    srcs = [
        "archive_test.go",
        "likes_test.go",
        "service_test.go",
        "subject_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/activity/conf",
        "//app/admin/main/activity/model",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
