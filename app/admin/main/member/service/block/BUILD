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
        "//app/admin/main/member/conf",
        "//app/admin/main/member/model/block",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "block",
    srcs = [
        "audit_log.go",
        "block.go",
        "msg.go",
        "notify.go",
        "rpc.go",
        "service.go",
    ],
    importpath = "go-common/app/admin/main/member/service/block",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/member/conf",
        "//app/admin/main/member/dao/block",
        "//app/admin/main/member/model/block",
        "//app/service/main/account/api",
        "//app/service/main/figure/model",
        "//app/service/main/figure/rpc/client",
        "//app/service/main/spy/model",
        "//app/service/main/spy/rpc/client",
        "//library/database/sql",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus",
        "//library/queue/databus/report",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
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
