package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "gitlab",
    srcs = ["gitlab.go"],
    importpath = "go-common/app/admin/ep/saga/service/gitlab",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/pkg/errors",
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
    name = "gitlab_test",
    srcs = ["gitlab_test.go"],
    embed = [":gitlab"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/ep/saga/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
