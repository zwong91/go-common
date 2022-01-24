package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "operation_test",
    srcs = ["promote_test.go"],
    embed = [":operation"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/web-show/conf",
        "//app/interface/main/web-show/model/operation",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "operation",
    srcs = [
        "notice.go",
        "promote.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/web-show/service/operation",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/web-show/conf",
        "//app/interface/main/web-show/dao/operation",
        "//app/interface/main/web-show/model/operation",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//library/ecode",
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
