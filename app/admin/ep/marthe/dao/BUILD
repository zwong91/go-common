package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "bugly.go",
        "dao.go",
        "mail.go",
        "mysql_bugly_batch_run.go",
        "mysql_bugly_cookie.go",
        "mysql_bugly_issue.go",
        "mysql_bugly_project.go",
        "mysql_bugly_version.go",
        "mysql_contact_info.go",
        "mysql_schedule_task.go",
        "mysql_tapd_bug_priority_conf.go",
        "mysql_tapd_bug_record.go",
        "mysql_tapd_bug_template.go",
        "mysql_tapd_bug_version_template.go",
        "mysql_user.go",
        "tapd.go",
        "wechat.go",
    ],
    importpath = "go-common/app/admin/ep/marthe/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/ep/marthe/conf",
        "//app/admin/ep/marthe/model",
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

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "mysql_bugly_batch_run_test.go",
        "mysql_bugly_cookie_test.go",
        "mysql_bugly_issue_test.go",
        "mysql_bugly_version_test.go",
        "mysql_schedule_task_test.go",
        "mysql_tapd_bug_priority_confg_test.go",
        "mysql_tapd_bug_record_test.go",
        "mysql_tapd_bug_template_test.go",
        "mysql_tapd_bug_version_template_test.go",
        "mysql_user_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/ep/marthe/conf",
        "//app/admin/ep/marthe/model",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/satori/go.uuid:go_uuid",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)
