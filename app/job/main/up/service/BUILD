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
        "service_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/up/conf",
        "//app/job/main/up/dao/upcrm",
        "//app/job/main/up/model/archivemodel",
        "//app/job/main/up/model/signmodel",
    ],
)

go_library(
    name = "service",
    srcs = [
        "archive.go",
        "service.go",
        "sign_due_job.go",
        "sign_state_job.go",
        "sign_task_job.go",
        "sign_update_tid_job.go",
    ],
    importpath = "go-common/app/job/main/up/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/util",
        "//app/admin/main/up/util/databusutil",
        "//app/admin/main/up/util/mathutil",
        "//app/interface/main/mcn/tool/worker",
        "//app/job/main/up/conf",
        "//app/job/main/up/dao/account",
        "//app/job/main/up/dao/email",
        "//app/job/main/up/dao/upcrm",
        "//app/job/main/up/model",
        "//app/job/main/up/model/archivemodel",
        "//app/job/main/up/model/signmodel",
        "//app/job/main/up/model/upcrmmodel",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/up/api/v1:api",
        "//app/service/main/upcredit/mathutil",
        "//library/log",
        "//library/queue/databus",
        "//library/time",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/robfig/cron",
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
