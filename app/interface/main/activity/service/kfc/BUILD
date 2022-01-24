package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "kfc_test",
    srcs = ["kfc_test.go"],
    embed = [":kfc"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/activity/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "kfc",
    srcs = ["kfc.go"],
    importpath = "go-common/app/interface/main/activity/service/kfc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/dao/kfc",
        "//app/interface/main/activity/model/kfc",
        "//library/ecode",
        "//library/log",
        "//library/sync/pipeline/fanout",
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
