package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["service_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/upcredit/conf",
        "//vendor/github.com/davecgh/go-spew/spew",
    ],
)

go_library(
    name = "service",
    srcs = [
        "business_score.go",
        "calculate_service.go",
        "content_ext.go",
        "creditscore.go",
        "databus.go",
        "logcredit.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/upcredit/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/upcredit/common/election",
        "//app/service/main/upcredit/common/fsm",
        "//app/service/main/upcredit/conf",
        "//app/service/main/upcredit/dao/upcrmdao",
        "//app/service/main/upcredit/mathutil",
        "//app/service/main/upcredit/model/calculator",
        "//app/service/main/upcredit/model/canal",
        "//app/service/main/upcredit/model/upcrmmodel",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/render",
        "//library/queue/databus",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/siddontang/go-mysql/mysql",
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
