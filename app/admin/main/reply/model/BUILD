package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "model_test",
    srcs = ["report_test.go"],
    embed = [":model"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "model",
    srcs = [
        "action.go",
        "admin.go",
        "business.go",
        "common.go",
        "config.go",
        "csv.go",
        "emoji.go",
        "monitor.go",
        "moral.go",
        "notice.go",
        "reply.go",
        "report.go",
        "sql.go",
        "subject.go",
    ],
    importpath = "go-common/app/admin/main/reply/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/relation/model",
        "//library/ecode",
        "//library/time",
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
