package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "gorpc_test",
    srcs = ["rpc_test.go"],
    embed = [":gorpc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/member/conf",
        "//app/service/main/member/model",
        "//app/service/main/member/service",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
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
    importpath = "go-common/app/service/main/member/server/gorpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/member/conf",
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//app/service/main/member/service",
        "//app/service/main/member/service/block",
        "//library/net/rpc",
        "//library/net/rpc/context",
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
