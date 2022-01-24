package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "cal_diff.go",
        "push.go",
        "service.go",
        "upbfs.go",
    ],
    importpath = "go-common/app/job/main/appstatic/service",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/appstatic/conf",
        "//app/job/main/appstatic/dao/caldiff",
        "//app/job/main/appstatic/dao/push",
        "//app/job/main/appstatic/model",
        "//library/log",
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
