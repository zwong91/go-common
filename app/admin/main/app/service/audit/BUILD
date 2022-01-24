package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "audit_test",
    srcs = ["audit_test.go"],
    embed = [":audit"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/app/conf",
        "//app/admin/main/app/model/audit",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "audit",
    srcs = ["audit.go"],
    importpath = "go-common/app/admin/main/app/service/audit",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/app/conf",
        "//app/admin/main/app/dao/audit",
        "//app/admin/main/app/model/audit",
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
