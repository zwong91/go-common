package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "service.go",
        "sns.go",
    ],
    importpath = "go-common/app/service/main/passport-sns/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/passport-sns/api",
        "//app/service/main/passport-sns/conf",
        "//app/service/main/passport-sns/dao",
        "//app/service/main/passport-sns/model",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/sync/pipeline/fanout",
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
