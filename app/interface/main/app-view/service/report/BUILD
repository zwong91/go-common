package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "report_test",
    srcs = ["report_test.go"],
    embed = [":report"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-view/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "report",
    srcs = ["report.go"],
    importpath = "go-common/app/interface/main/app-view/service/report",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-view/conf",
        "//app/interface/main/app-view/dao/archive",
        "//app/interface/main/app-view/dao/elec",
        "//app/interface/main/app-view/dao/report",
        "//app/interface/main/app-view/model",
        "//app/interface/main/app-view/model/elec",
        "//app/interface/main/app-view/model/report",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//library/ecode",
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
