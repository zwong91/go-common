package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "bugly.go",
        "bugly_batch_run.go",
        "mail.go",
        "service.go",
        "tapd.go",
        "user.go",
    ],
    importpath = "go-common/app/admin/ep/marthe/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/ep/marthe/conf",
        "//app/admin/ep/marthe/dao",
        "//app/admin/ep/marthe/model",
        "//library/ecode",
        "//library/log",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/robfig/cron",
        "//vendor/github.com/satori/go.uuid:go_uuid",
        "//vendor/gopkg.in/gomail.v2:gomail_v2",
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
