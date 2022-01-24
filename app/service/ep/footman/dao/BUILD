package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "bugly.go",
        "dao.go",
        "mail.go",
        "mysql_bug_template.go",
        "mysql_bugly.go",
        "tapd.go",
    ],
    importpath = "go-common/app/service/ep/footman/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/ep/footman/conf",
        "//app/service/ep/footman/model",
        "//library/cache/memcache",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/jinzhu/gorm",
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
