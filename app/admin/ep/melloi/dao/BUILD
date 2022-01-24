package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "order_admin_test.go",
        "order_test.go",
        "rank_test.go",
        "scene_test.go",
        "user_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/ep/melloi/conf",
        "//app/admin/ep/melloi/model",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "apply.go",
        "bfs.go",
        "clientmoni.go",
        "cluster.go",
        "comment.go",
        "dao.go",
        "dapper.go",
        "job.go",
        "label.go",
        "mail.go",
        "mysql_grpc.go",
        "mysql_grpc_snap.go",
        "order.go",
        "order_admin.go",
        "order_report.go",
        "ptestjob.go",
        "rank.go",
        "report.go",
        "reportgraph.go",
        "scene.go",
        "script.go",
        "tree.go",
        "user.go",
        "wechat.go",
    ],
    importpath = "go-common/app/admin/ep/melloi/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/ep/melloi/conf",
        "//app/admin/ep/melloi/model",
        "//library/cache/memcache",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
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
