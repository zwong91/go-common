package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "activity_test.go",
        "business_test.go",
        "challenge_test.go",
        "platform_test.go",
        "service_test.go",
        "wlog_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/workflow/model",
        "//app/admin/main/workflow/model/param",
        "//app/admin/main/workflow/model/search",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "activity.go",
        "business.go",
        "callback.go",
        "challenge.go",
        "event.go",
        "group.go",
        "meta.go",
        "platform.go",
        "service.go",
        "tag.go",
        "task.go",
        "wlog.go",
    ],
    importpath = "go-common/app/admin/main/workflow/service",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/credit/model/blocked",
        "//app/admin/main/workflow/dao",
        "//app/admin/main/workflow/model",
        "//app/admin/main/workflow/model/param",
        "//app/admin/main/workflow/model/search",
        "//app/job/main/credit/model",
        "//library/conf/paladin",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/queue/databus/report",
        "//library/time",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/pkg/errors",
        "//vendor/golang.org/x/sync/errgroup",
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
