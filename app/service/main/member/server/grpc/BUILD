package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "grpc",
    srcs = [
        "block.go",
        "exp.go",
        "member.go",
        "moral.go",
        "property_review.go",
        "realname.go",
    ],
    importpath = "go-common/app/service/main/member/server/grpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/member/api",
        "//app/service/main/member/model",
        "//app/service/main/member/service",
        "//app/service/main/member/service/block",
        "//library/net/rpc/warden",
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
