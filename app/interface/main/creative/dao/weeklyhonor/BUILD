package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "weeklyhonor",
    srcs = [
        "dao.go",
        "hbase.go",
        "memcache.go",
        "mysql.go",
        "rpc.go",
    ],
    importpath = "go-common/app/interface/main/creative/dao/weeklyhonor",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/util/hbaseutil",
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/weeklyhonor",
        "//app/service/main/up/api/v1:api",
        "//library/cache/memcache",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//vendor/github.com/tsuna/gohbase/hrpc",
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
    name = "weeklyhonor_test",
    srcs = [
        "dao_test.go",
        "hbase_test.go",
        "memcache_test.go",
        "mysql_test.go",
        "rpc_test.go",
    ],
    embed = [":weeklyhonor"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/weeklyhonor",
        "//app/service/main/up/api/v1:api",
        "//library/database/hbase.v2:hbase_v2",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/github.com/tsuna/gohbase/hrpc",
    ],
)
