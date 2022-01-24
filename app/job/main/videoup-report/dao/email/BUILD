package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "email_test",
    srcs = [
        "dao_test.go",
        "email_template_test.go",
        "email_test.go",
        "redis_list_test.go",
    ],
    embed = [":email"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/videoup-report/conf",
        "//app/job/main/videoup-report/model/archive",
        "//app/job/main/videoup-report/model/email",
        "//library/cache/redis",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "email",
    srcs = [
        "dao.go",
        "email.go",
        "email_template.go",
        "redis_list.go",
    ],
    importpath = "go-common/app/job/main/videoup-report/dao/email",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/videoup-report/conf",
        "//app/job/main/videoup-report/model/email",
        "//library/cache/redis",
        "//library/log",
        "//library/queue/databus/report",
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
