package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "advance_test.go",
        "assist_test.go",
        "dm_test.go",
        "filter_test.go",
        "protect_test.go",
        "report_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/dm/conf",
        "//app/interface/main/dm/model",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "advance.go",
        "archive.go",
        "assist.go",
        "dm.go",
        "filter.go",
        "protect.go",
        "report.go",
        "service.go",
        "subject.go",
    ],
    importpath = "go-common/app/interface/main/dm/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/dm/conf",
        "//app/interface/main/dm/dao",
        "//app/interface/main/dm/model",
        "//app/interface/main/dm2/model",
        "//app/interface/main/dm2/model/oplog",
        "//app/interface/main/dm2/rpc/client",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/assist/model/assist",
        "//app/service/main/assist/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/xstr",
        "//vendor/github.com/zhenjl/cityhash",
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
