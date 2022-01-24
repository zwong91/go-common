package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "academy_test",
    srcs = ["service_test.go"],
    embed = [":academy"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/academy",
        "//app/interface/main/creative/service",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "academy",
    srcs = [
        "academy.go",
        "h5.go",
        "service.go",
        "skill.go",
    ],
    importpath = "go-common/app/interface/main/creative/service/academy",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/academy",
        "//app/interface/main/creative/dao/archive",
        "//app/interface/main/creative/dao/article",
        "//app/interface/main/creative/dao/resource",
        "//app/interface/main/creative/dao/tool",
        "//app/interface/main/creative/model/academy",
        "//app/interface/main/creative/model/newcomer",
        "//app/interface/main/creative/service",
        "//app/interface/openplatform/article/model",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup",
        "//library/time",
        "//vendor/github.com/davecgh/go-spew/spew",
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
