package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "audit.go",
        "bilihub.go",
        "cluster.go",
        "dashboard.go",
        "devicefarm.go",
        "image.go",
        "machine.go",
        "machineV2.go",
        "machine_log.go",
        "mail.go",
        "node.go",
        "service.go",
        "task.go",
        "tree.go",
        "user.go",
    ],
    importpath = "go-common/app/admin/ep/merlin/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/ep/merlin/conf",
        "//app/admin/ep/merlin/dao",
        "//app/admin/ep/merlin/model",
        "//library/ecode",
        "//library/log",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/robfig/cron",
        "//vendor/github.com/satori/go.uuid:go_uuid",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "service_test",
    srcs = [
        "machine_test.go",
        "service_test.go",
        "tree_test.go",
        "user_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/ep/merlin/conf",
        "//app/admin/ep/merlin/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)
