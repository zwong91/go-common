package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "hook_url.go",
        "mysql_event_log.go",
        "mysql_hook_url.go",
        "mysql_url_event.go",
    ],
    importpath = "go-common/app/admin/ep/tapd/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/ep/tapd/conf",
        "//app/admin/ep/tapd/model",
        "//library/cache/memcache",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/jinzhu/gorm",
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
