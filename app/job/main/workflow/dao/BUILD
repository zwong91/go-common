load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "appeal.go",
        "buss_attr.go",
        "chall.go",
        "dao.go",
        "notify.go",
        "redis.go",
        "search.go",
    ],
    importpath = "go-common/app/job/main/workflow/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/workflow/model/param",
        "//app/job/main/workflow/conf",
        "//app/job/main/workflow/model",
        "//library/cache/redis",
        "//library/database/elastic",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/queue/databus/report",
        "//vendor/github.com/jinzhu/gorm",
    ],
)

go_test(
    name = "dao_test",
    srcs = [
        "appeal_test.go",
        "buss_attr_test.go",
        "chall_test.go",
        "dao_test.go",
        "notify_test.go",
        "redis_test.go",
        "search_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/workflow/model/param",
        "//app/job/main/workflow/conf",
        "//app/job/main/workflow/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
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
