package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "data_test",
    srcs = [
        "admin_test.go",
        "dao_test.go",
        "reply_test.go",
    ],
    embed = [":data"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/videoup-report/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "data",
    srcs = [
        "admin.go",
        "dao.go",
        "profit.go",
        "reply.go",
        "uper.go",
    ],
    importpath = "go-common/app/job/main/videoup-report/dao/data",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/videoup-report/conf",
        "//app/job/main/videoup-report/model/archive",
        "//app/job/main/videoup-report/model/data",
        "//app/job/main/videoup-report/model/manager",
        "//app/job/main/videoup-report/model/monitor",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
