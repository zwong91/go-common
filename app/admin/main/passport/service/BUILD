package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "aes.go",
        "service.go",
        "user_act_log.go",
    ],
    importpath = "go-common/app/admin/main/passport/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/passport/conf",
        "//app/admin/main/passport/dao",
        "//app/admin/main/passport/model",
        "//library/database/elastic",
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

go_test(
    name = "service_test",
    srcs = ["user_act_log_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/passport/conf",
        "//app/admin/main/passport/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
