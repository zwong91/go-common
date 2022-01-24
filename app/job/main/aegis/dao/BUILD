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
        "task_db.go",
        "task_http.go",
        "task_mc.go",
        "task_orm.go",
        "task_redis.go",
        "task_report.go",
        "task_rpc.go",
    ],
    importpath = "go-common/app/job/main/aegis/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/aegis/conf",
        "//app/job/main/aegis/model",
        "//app/service/main/relation/model",
        "//app/service/main/relation/rpc/client",
        "//app/service/main/up/api/v1:api",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/orm",
        "//library/database/sql",
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
    srcs = [
        ":package-srcs",
        "//app/job/main/aegis/dao/email:all-srcs",
        "//app/job/main/aegis/dao/monitor:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "task_db_test.go",
        "task_mc_test.go",
        "task_orm_test.go",
        "task_redis_test.go",
        "task_rpc_test.go",
        "task_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/aegis/conf",
        "//app/job/main/aegis/model",
        "//library/sync/errgroup",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
