package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dataflow",
    srcs = ["service.go"],
    importpath = "go-common/app/interface/main/app-interface/service/dataflow",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//library/log",
        "//library/log/infoc",
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
    name = "dataflow_test",
    srcs = ["service_test.go"],
    embed = [":dataflow"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
