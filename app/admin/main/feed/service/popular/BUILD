load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "popular",
    srcs = [
        "event_topic.go",
        "service.go",
        "stars.go",
    ],
    importpath = "go-common/app/admin/main/feed/service/popular",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/feed/conf",
        "//app/admin/main/feed/dao/show",
        "//app/admin/main/feed/model/common",
        "//app/admin/main/feed/model/show",
        "//app/admin/main/feed/util",
        "//library/log",
        "//vendor/github.com/jinzhu/gorm",
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
