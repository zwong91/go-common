package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "mcnmodel",
    srcs = [
        "cmd.go",
        "creative.go",
        "mcncache.go",
        "reply.go",
        "req.go",
        "tables.go",
        "upcheck.go",
    ],
    importpath = "go-common/app/interface/main/mcn/model/mcnmodel",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/mcn/model",
        "//app/interface/main/mcn/model",
        "//app/interface/main/mcn/model/datamodel",
        "//app/interface/main/mcn/tool/validate",
        "//app/service/main/archive/api",
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

go_test(
    name = "mcnmodel_test",
    srcs = ["reply_test.go"],
    embed = [":mcnmodel"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/mcn/model",
        "//app/interface/main/mcn/model",
    ],
)
