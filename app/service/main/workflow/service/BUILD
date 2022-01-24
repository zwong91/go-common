package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "business_test.go",
        "challenge_test.go",
        "service_test.go",
        "tags_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/workflow/conf",
        "//app/service/main/workflow/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "business.go",
        "challenge.go",
        "event.go",
        "group.go",
        "service.go",
        "sobot.go",
        "tags.go",
    ],
    importpath = "go-common/app/service/main/workflow/service",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/workflow/conf",
        "//app/service/main/workflow/dao",
        "//app/service/main/workflow/dao/sobot",
        "//app/service/main/workflow/model",
        "//app/service/main/workflow/model/sobot",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/jinzhu/gorm",
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
