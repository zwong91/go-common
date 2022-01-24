package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "bilihub.go",
        "dao.go",
        "dashboard.go",
        "devicefarm.go",
        "docker.go",
        "mail.go",
        "mysql_application_record.go",
        "mysql_dashboard.go",
        "mysql_hubimage_confs.go",
        "mysql_hubimage_logs.go",
        "mysql_image.go",
        "mysql_machine.go",
        "mysql_machineV2.go",
        "mysql_machine_log.go",
        "mysql_machine_node.go",
        "mysql_machine_package.go",
        "mysql_mail_log.go",
        "mysql_mobile_machine.go",
        "mysql_mobile_machine_error_log.go",
        "mysql_mobile_machine_log.go",
        "mysql_mobile_sync_log.go",
        "mysql_snapshot_record.go",
        "mysql_task.go",
        "mysql_user.go",
        "paas.go",
        "tree.go",
        "wechat.go",
    ],
    importpath = "go-common/app/admin/ep/merlin/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/ep/merlin/conf",
        "//app/admin/ep/merlin/model",
        "//library/cache/memcache",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/sync/pipeline/fanout",
        "//vendor/github.com/docker/docker/api/types",
        "//vendor/github.com/docker/docker/api/types/filters",
        "//vendor/github.com/docker/docker/client",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/pkg/errors",
        "//vendor/gopkg.in/gomail.v2:gomail_v2",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
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
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "mail_test.go",
        "mysql_image_test.go",
        "mysql_machine_log_test.go",
        "mysql_machine_test.go",
        "mysql_mail_log_test.go",
        "mysql_task_test.go",
        "mysql_user_test.go",
        "paas_test.go",
        "tree_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/ep/merlin/conf",
        "//app/admin/ep/merlin/model",
        "//library/log",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/gomail.v2:gomail_v2",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)
