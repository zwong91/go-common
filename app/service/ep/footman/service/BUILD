package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "bigdata.go",
        "bugly.go",
        "bugly2tapd.go",
        "excel.go",
        "mail.go",
        "report.go",
        "service.go",
        "tapd.go",
        "util.go",
    ],
    importpath = "go-common/app/service/ep/footman/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/ep/footman/conf",
        "//app/service/ep/footman/dao",
        "//app/service/ep/footman/model",
        "//library/cache",
        "//library/log",
        "//vendor/github.com/360EntSecGroup-Skylar/excelize",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
