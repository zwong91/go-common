package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "common",
    srcs = [
        "cardpreview.go",
        "cardtype.go",
        "common.go",
        "log.go",
    ],
    importpath = "go-common/app/admin/main/feed/service/common",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/feed/conf",
        "//app/admin/main/feed/dao/account",
        "//app/admin/main/feed/dao/archive",
        "//app/admin/main/feed/dao/pgc",
        "//app/admin/main/feed/dao/show",
        "//app/admin/main/feed/model/common",
        "//app/admin/main/feed/model/show",
        "//app/admin/main/search/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/openplatform/pgc-season/api/grpc/season/v1:season",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
