package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "bfs_test.go",
        "dao_test.go",
        "dialog_test.go",
        "jointly_test.go",
        "memcache_test.go",
        "month_test.go",
        "mysql_test.go",
        "order_test.go",
        "panel_test.go",
        "platform_test.go",
        "privilege_test.go",
        "push_test.go",
        "resouce_test.go",
        "send_test.go",
        "tips_test.go",
        "vip_test.go",
        "welfare_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/vip/conf",
        "//app/admin/main/vip/model",
        "//library/database/sql",
        "//library/log",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "bfs.go",
        "dao.go",
        "dialog.go",
        "jointly.go",
        "memcache.go",
        "month.go",
        "mysql.go",
        "order.go",
        "panel.go",
        "platform.go",
        "privilege.go",
        "push.go",
        "resouce.go",
        "send.go",
        "tips.go",
        "vip.go",
        "welfare.go",
    ],
    importpath = "go-common/app/admin/main/vip/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/vip/conf",
        "//app/admin/main/vip/model",
        "//library/cache/memcache",
        "//library/database/orm",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//library/time",
        "//vendor/github.com/google/uuid",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/pkg/errors",
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
