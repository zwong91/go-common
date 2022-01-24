package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "audit_log.go",
        "mcn.go",
        "msg.go",
        "recommend.go",
        "service.go",
        "statistics.go",
        "upload.go",
    ],
    importpath = "go-common/app/admin/main/mcn/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/mcn/conf",
        "//app/admin/main/mcn/dao/bfs",
        "//app/admin/main/mcn/dao/msg",
        "//app/admin/main/mcn/dao/up",
        "//app/admin/main/mcn/dao/videoup",
        "//app/admin/main/mcn/model",
        "//app/interface/main/mcn/model/datamodel",
        "//app/interface/main/mcn/tool/worker",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/member/api",
        "//app/service/main/member/model/block",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus/report",
        "//library/sync/errgroup",
        "//library/time",
        "//library/xstr",
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

go_test(
    name = "service_test",
    srcs = [
        "audit_log_test.go",
        "mcn_test.go",
        "msg_test.go",
        "recommend_test.go",
        "service_test.go",
        "statistics_test.go",
        "upload_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/mcn/conf",
        "//app/admin/main/mcn/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
