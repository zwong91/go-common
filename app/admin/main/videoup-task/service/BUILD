package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "search_test.go",
        "service_test.go",
        "video_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/videoup-task/conf",
        "//app/admin/main/videoup-task/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "consumer.go",
        "extra_func.go",
        "log.go",
        "oper.go",
        "qa_video.go",
        "review.go",
        "search.go",
        "service.go",
        "task.go",
        "task_dispatch.go",
        "task_report.go",
        "utils.go",
        "video.go",
        "weight.go",
    ],
    importpath = "go-common/app/admin/main/videoup-task/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/videoup-task/conf",
        "//app/admin/main/videoup-task/dao",
        "//app/admin/main/videoup-task/model",
        "//app/admin/main/videoup/model/archive",
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//app/service/main/up/api/v1:api",
        "//library/database/elastic",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/queue/databus/report",
        "//library/sync/errgroup",
        "//library/xstr",
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
