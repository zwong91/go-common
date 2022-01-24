package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "block_test",
    srcs = ["service_test.go"],
    embed = [":block"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/member/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "block",
    srcs = [
        "audit_log.go",
        "blackhouse.go",
        "limit.go",
        "msg.go",
        "service.go",
    ],
    importpath = "go-common/app/job/main/member/service/block",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/member/conf",
        "//app/job/main/member/dao/block",
        "//app/job/main/member/model/block",
        "//library/cache",
        "//library/database/sql",
        "//library/log",
        "//library/queue/databus",
        "//library/queue/databus/report",
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
