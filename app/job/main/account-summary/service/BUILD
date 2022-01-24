package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "block.go",
        "initial.go",
        "member.go",
        "passport.go",
        "relation.go",
        "service.go",
        "summary.go",
        "sync.go",
    ],
    importpath = "go-common/app/job/main/account-summary/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/account-summary/conf",
        "//app/job/main/account-summary/dao",
        "//app/job/main/account-summary/model",
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//app/service/main/relation/model",
        "//library/log",
        "//library/queue/databus",
        "//library/time",
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
