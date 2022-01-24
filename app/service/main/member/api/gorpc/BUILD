package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "gorpc",
    srcs = [
        "block.go",
        "member.go",
        "moral.go",
        "property_review.go",
        "realname.go",
        "rpc.go",
    ],
    importpath = "go-common/app/service/main/member/api/gorpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//library/ecode",
        "//library/net/rpc",
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
