package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "command",
    srcs = [
        "command.go",
        "contributors.go",
        "merge.go",
        "review.go",
    ],
    importpath = "go-common/app/tool/saga/service/command",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/saga/conf",
        "//app/tool/saga/dao",
        "//app/tool/saga/model",
        "//app/tool/saga/service/gitlab",
        "//app/tool/saga/service/notification",
        "//library/log",
        "//vendor/github.com/xanzy/go-gitlab",
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

go_test(
    name = "command_test",
    srcs = [
        "command_test.go",
        "contributors_test.go",
        "review_test.go",
    ],
    embed = [":command"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/tool/saga/conf",
        "//app/tool/saga/dao",
        "//app/tool/saga/model",
        "//app/tool/saga/service/gitlab",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
