package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "comment.go",
        "contributors.go",
        "mr.go",
        "pipeline.go",
        "service.go",
        "wechat.go",
    ],
    importpath = "go-common/app/tool/saga/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/saga/conf",
        "//app/tool/saga/dao",
        "//app/tool/saga/model",
        "//app/tool/saga/service/command",
        "//app/tool/saga/service/gitlab",
        "//app/tool/saga/service/mail",
        "//app/tool/saga/service/notification",
        "//app/tool/saga/service/wechat",
        "//library/log",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/robfig/cron",
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
        "//app/tool/saga/service/command:all-srcs",
        "//app/tool/saga/service/gitlab:all-srcs",
        "//app/tool/saga/service/mail:all-srcs",
        "//app/tool/saga/service/notification:all-srcs",
        "//app/tool/saga/service/wechat:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "service_test",
    srcs = [
        "comment_test.go",
        "contributors_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/tool/saga/conf",
        "//app/tool/saga/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
