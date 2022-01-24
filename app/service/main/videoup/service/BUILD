package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "archive_test.go",
        "audit_log_test.go",
        "flow_test.go",
        "msg_test.go",
        "pgc_test.go",
        "porder_test.go",
        "service_test.go",
        "staff_test.go",
        "ugc_test.go",
        "util_test.go",
        "video_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/videoup/conf",
        "//app/service/main/videoup/model/archive",
        "//library/queue/databus/report",
        "//vendor/github.com/davecgh/go-spew/spew",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "archive.go",
        "audit_log.go",
        "databus.go",
        "dede.go",
        "extra_func.go",
        "flow.go",
        "monitor.go",
        "msg.go",
        "netsafe.go",
        "pgc.go",
        "porder.go",
        "report.go",
        "service.go",
        "staff.go",
        "tx.go",
        "ugc.go",
        "util.go",
        "video.go",
    ],
    importpath = "go-common/app/service/main/videoup/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/mcn/tool/worker",
        "//app/service/main/account/api",
        "//app/service/main/videoup/conf",
        "//app/service/main/videoup/dao/agent",
        "//app/service/main/videoup/dao/archive",
        "//app/service/main/videoup/dao/bgm",
        "//app/service/main/videoup/dao/databus",
        "//app/service/main/videoup/dao/dede",
        "//app/service/main/videoup/dao/manager",
        "//app/service/main/videoup/dao/monitor",
        "//app/service/main/videoup/dao/msg",
        "//app/service/main/videoup/dao/relation",
        "//app/service/main/videoup/dao/ups",
        "//app/service/main/videoup/model/archive",
        "//app/service/main/videoup/model/dede",
        "//app/service/main/videoup/model/message",
        "//app/service/main/videoup/model/prom",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus",
        "//library/queue/databus/report",
        "//library/stat/prom",
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
