package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "apply_test.go",
        "grpc_snap_test.go",
        "grpc_test.go",
        "label_test.go",
        "mail_test.go",
        "order_admin_test.go",
        "order_test.go",
        "rank_test.go",
        "report_test.go",
        "scene_test.go",
        "script_test.go",
        "service_test.go",
        "user_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/ep/melloi/conf",
        "//app/admin/ep/melloi/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "apply.go",
        "bfs.go",
        "clientmoni.go",
        "cluster.go",
        "comment.go",
        "dapper.go",
        "file.go",
        "grpc.go",
        "grpc_snap.go",
        "job.go",
        "label.go",
        "mail.go",
        "order.go",
        "order_admin.go",
        "order_report.go",
        "ptest.go",
        "rank.go",
        "report.go",
        "scene.go",
        "script.go",
        "script_snap.go",
        "service.go",
        "tools.go",
        "tree.go",
        "user.go",
        "wechat.go",
    ],
    importpath = "go-common/app/admin/ep/melloi/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/ep/melloi/conf",
        "//app/admin/ep/melloi/dao",
        "//app/admin/ep/melloi/model",
        "//app/admin/ep/melloi/service/proto",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/robfig/cron",
        "//vendor/gopkg.in/gomail.v2:gomail_v2",
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
        "//app/admin/ep/melloi/service/proto:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
