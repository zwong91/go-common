package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "action_test.go",
        "report_test.go",
        "resource_test.go",
        "service_test.go",
        "sub_sort_test.go",
        "sub_test.go",
        "tag_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/tag/conf",
        "//library/cache/redis",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "action.go",
        "channel.go",
        "check.go",
        "report.go",
        "resource.go",
        "service.go",
        "sub.go",
        "sub_sort.go",
        "tag.go",
    ],
    importpath = "go-common/app/service/main/tag/service",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/tag/api",
        "//app/service/main/tag/conf",
        "//app/service/main/tag/dao",
        "//app/service/main/tag/model",
        "//library/cache",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/time",
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
