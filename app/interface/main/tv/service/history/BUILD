package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "history",
    srcs = [
        "combine.go",
        "history.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/tv/service/history",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/history/model",
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/dao/archive",
        "//app/interface/main/tv/dao/cms",
        "//app/interface/main/tv/dao/history",
        "//app/interface/main/tv/model",
        "//app/interface/main/tv/model/history",
        "//library/log",
        "//library/sync/errgroup",
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
    name = "history_test",
    srcs = ["service_test.go"],
    embed = [":history"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tv/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
