package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "model",
    srcs = [
        "history.go",
        "rpc.go",
        "toview.go",
    ],
    importpath = "go-common/app/interface/main/history/model",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/model/archive",
        "//app/service/main/history/model",
        "//library/ecode",
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
