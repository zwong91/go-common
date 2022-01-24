load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "channel",
    srcs = [
        "card_setup.go",
        "tab.go",
    ],
    importpath = "go-common/app/admin/main/feed/service/channel",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/feed/conf",
        "//app/admin/main/feed/dao/account",
        "//app/admin/main/feed/dao/archive",
        "//app/admin/main/feed/dao/pgc",
        "//app/admin/main/feed/dao/show",
        "//app/admin/main/feed/model/channel",
        "//app/admin/main/feed/model/common",
        "//app/admin/main/feed/model/show",
        "//app/admin/main/feed/util",
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
