package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "audit_test",
    srcs = ["service_test.go"],
    embed = [":audit"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//app/interface/main/tv/conf"],
)

go_library(
    name = "audit",
    srcs = [
        "audit.go",
        "service.go",
        "transcode.go",
    ],
    importpath = "go-common/app/interface/main/tv/service/audit",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/dao/audit",
        "//app/interface/main/tv/dao/cms",
        "//app/interface/main/tv/model",
        "//library/database/sql",
        "//library/ecode",
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
