package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "region_test",
    srcs = [
        "infoc_test.go",
        "region_test.go",
        "service_test.go",
    ],
    embed = [":region"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-feed/conf",
        "//app/interface/main/app-feed/model/tag",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "region",
    srcs = [
        "infoc.go",
        "region.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/app-feed/service/region",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-feed/conf",
        "//app/interface/main/app-feed/dao/tag",
        "//app/interface/main/app-feed/model",
        "//app/interface/main/app-feed/model/tag",
        "//library/ecode",
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
