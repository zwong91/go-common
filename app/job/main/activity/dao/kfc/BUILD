package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "kfc_test",
    srcs = [
        "dao_test.go",
        "kfc_test.go",
    ],
    embed = [":kfc"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/activity/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "kfc",
    srcs = [
        "dao.go",
        "kfc.go",
    ],
    importpath = "go-common/app/job/main/activity/dao/kfc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/activity/conf",
        "//library/ecode",
        "//library/net/http/blademaster",
        "//library/net/metadata",
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
