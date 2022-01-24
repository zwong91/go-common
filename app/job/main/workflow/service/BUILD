load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "appeal.go",
        "expire.go",
        "notify.go",
        "queue.go",
        "service.go",
        "workflow.go",
    ],
    importpath = "go-common/app/job/main/workflow/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/workflow/model/param",
        "//app/job/main/workflow/conf",
        "//app/job/main/workflow/dao",
        "//app/job/main/workflow/model",
        "//app/service/main/workflow/model",
        "//library/database/elastic",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup.v2:errgroup_v2",
        "//library/sync/pipeline/fanout",
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
    srcs = ["service_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/workflow/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
