package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "relation.go",
        "service.go",
    ],
    importpath = "go-common/app/job/main/relation-cache/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/relation-cache/conf",
        "//app/job/main/relation-cache/dao",
        "//app/job/main/relation-cache/model",
        "//app/service/main/relation/model",
        "//library/log",
        "//library/queue/databus",
        "//library/time",
        "//vendor/github.com/pkg/errors",
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
