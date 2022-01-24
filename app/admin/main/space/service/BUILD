package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "account.go",
        "blacklist.go",
        "notice.go",
        "service.go",
    ],
    importpath = "go-common/app/admin/main/space/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/space/conf",
        "//app/admin/main/space/dao",
        "//app/admin/main/space/model",
        "//app/service/main/relation/model",
        "//app/service/main/relation/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus/report",
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
