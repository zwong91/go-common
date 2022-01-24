package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "appeal.go",
        "data_check.go",
        "send_email.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/account-recovery/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/account-recovery/conf",
        "//app/service/main/account-recovery/dao",
        "//app/service/main/account-recovery/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus",
        "//library/sync/errgroup",
        "//library/time",
        "//vendor/github.com/microcosm-cc/bluemonday",
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
    srcs = [
        "appeal_test.go",
        "data_check_test.go",
        "send_email_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/account-recovery/conf",
        "//app/service/main/account-recovery/model",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
