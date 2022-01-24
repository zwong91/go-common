package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "http_test",
    srcs = [
        "adv_test.go",
        "filter_test.go",
        "http_test.go",
    ],
    embed = [":http"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//app/admin/main/dm/conf"],
)

go_library(
    name = "http",
    srcs = [
        "adv.go",
        "dm.go",
        "filter.go",
        "http.go",
        "mask.go",
        "monitor.go",
        "report.go",
        "shield.go",
        "subject.go",
        "subtitle.go",
        "task.go",
        "transfer.go",
    ],
    importpath = "go-common/app/admin/main/dm/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/dm/conf",
        "//app/admin/main/dm/model",
        "//app/admin/main/dm/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/http/blademaster/middleware/verify",
        "//library/xstr",
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
