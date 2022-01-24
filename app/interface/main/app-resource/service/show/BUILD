package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "show_test",
    srcs = ["show_test.go"],
    embed = [":show"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-resource/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "show",
    srcs = [
        "audit.go",
        "cache.go",
        "service.go",
        "show.go",
    ],
    importpath = "go-common/app/interface/main/app-resource/service/show",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-resource/conf",
        "//app/interface/main/app-resource/dao/audit",
        "//app/interface/main/app-resource/dao/resource",
        "//app/interface/main/app-resource/dao/tab",
        "//app/interface/main/app-resource/model",
        "//app/interface/main/app-resource/model/abtest",
        "//app/interface/main/app-resource/model/show",
        "//app/interface/main/app-resource/model/tab",
        "//app/service/main/resource/model",
        "//library/ecode",
        "//library/log",
        "//vendor/github.com/dgryski/go-farm",
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
