package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "telecom_test",
    srcs = ["telecom_test.go"],
    embed = [":telecom"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-wall/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "telecom",
    srcs = [
        "service.go",
        "telecom.go",
    ],
    importpath = "go-common/app/interface/main/app-wall/service/telecom",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-wall/conf",
        "//app/interface/main/app-wall/dao/seq",
        "//app/interface/main/app-wall/dao/telecom",
        "//app/interface/main/app-wall/model/telecom",
        "//library/ecode",
        "//library/log",
        "//library/stat/prom",
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
