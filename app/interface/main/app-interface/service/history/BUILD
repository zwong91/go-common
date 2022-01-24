package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "history_test",
    srcs = [
        "history_test.go",
        "service_test.go",
    ],
    embed = [":history"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "history",
    srcs = [
        "history.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/app-interface/service/history",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/dao/history",
        "//app/interface/main/app-interface/dao/live",
        "//app/interface/main/app-interface/model",
        "//app/interface/main/app-interface/model/history",
        "//app/interface/main/app-interface/model/live",
        "//app/interface/main/history/model",
        "//app/interface/openplatform/article/model",
        "//app/service/main/archive/model/archive",
        "//library/log",
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
