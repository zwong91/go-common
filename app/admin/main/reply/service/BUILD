package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "action_test.go",
        "admin_test.go",
        "business_test.go",
        "config_test.go",
        "event_test.go",
        "moral_test.go",
        "notice_test.go",
        "reply_test.go",
        "report_test.go",
        "service_test.go",
        "subject_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/reply/conf",
        "//app/admin/main/reply/model",
        "//library/log",
        "//library/time",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "action.go",
        "admin.go",
        "business.go",
        "config.go",
        "emoji.go",
        "emoji_package.go",
        "event.go",
        "fold.go",
        "monitor.go",
        "moral.go",
        "notice.go",
        "reply.go",
        "report.go",
        "service.go",
        "subject.go",
    ],
    importpath = "go-common/app/admin/main/reply/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/reply/conf",
        "//app/admin/main/reply/dao",
        "//app/admin/main/reply/model",
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/rpc/client",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/relation/model",
        "//app/service/main/relation/rpc/client",
        "//app/service/main/thumbup/api",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/net/rpc/warden",
        "//library/queue/databus/report",
        "//library/sync/errgroup",
        "//library/sync/errgroup.v2:errgroup_v2",
        "//library/sync/pipeline/fanout",
        "//library/time",
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
