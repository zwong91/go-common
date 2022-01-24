package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "email.go",
        "memcache.go",
        "task.go",
        "task_log.go",
    ],
    importpath = "go-common/app/interface/main/laser/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/laser/conf",
        "//app/interface/main/laser/model",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/log",
        "//vendor/github.com/pkg/errors",
        "//vendor/gopkg.in/gomail.v2:gomail_v2",
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
