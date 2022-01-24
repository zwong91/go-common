package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "block",
    srcs = [
        "blackhouse.go",
        "dao.go",
        "model.go",
        "msg.go",
        "notify.go",
    ],
    importpath = "go-common/app/job/main/member/model/block",
    tags = ["automanaged"],
    deps = ["//library/time"],
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
