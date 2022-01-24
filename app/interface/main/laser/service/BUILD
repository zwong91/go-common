package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "mail.go",
        "service.go",
        "task.go",
    ],
    importpath = "go-common/app/interface/main/laser/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/laser/conf",
        "//app/interface/main/laser/dao",
        "//app/interface/main/laser/model",
        "//library/database/sql",
        "//library/log",
        "//library/stat/prom",
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
