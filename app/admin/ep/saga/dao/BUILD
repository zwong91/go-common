package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "db_project.go",
        "db_sync.go",
        "db_task.go",
        "db_wechat.go",
        "http.go",
        "mc.go",
        "redis.go",
    ],
    importpath = "go-common/app/admin/ep/saga/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/ep/saga/conf",
        "//app/admin/ep/saga/model",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "db_project_test.go",
        "db_sync_test.go",
        "db_wechat_test.go",
        "mc_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/ep/saga/conf",
        "//app/admin/ep/saga/model",
        "//app/admin/ep/saga/service/utils",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
