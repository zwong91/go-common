package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "account_test.go",
        "dao_test.go",
        "relation_log_test.go",
        "relation_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/relation/conf",
        "//library/database/hbase.v2:hbase_v2",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/github.com/tsuna/gohbase/hrpc",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "account.go",
        "dao.go",
        "relation.go",
        "relation_log.go",
    ],
    importpath = "go-common/app/admin/main/relation/dao",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/relation/conf",
        "//app/admin/main/relation/model",
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//app/service/main/relation/api",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/orm",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
        "//vendor/github.com/jinzhu/gorm",
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
