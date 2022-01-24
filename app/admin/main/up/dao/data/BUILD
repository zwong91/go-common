package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "data_test",
    srcs = [
        "api_test.go",
        "center_test.go",
        "dao_test.go",
        "field_test.go",
        "hbase_set2_test.go",
        "hbase_test.go",
        "phase_three_test.go",
        "utils_test.go",
    ],
    embed = [":data"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/up/conf",
        "//app/admin/main/up/model/datamodel",
        "//library/database/hbase.v2:hbase_v2",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/github.com/tsuna/gohbase/hrpc",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "data",
    srcs = [
        "api.go",
        "center.go",
        "dao.go",
        "field.go",
        "hbase.go",
        "hbase_set2.go",
        "phase_three.go",
        "utils.go",
    ],
    importpath = "go-common/app/admin/main/up/dao/data",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/conf",
        "//app/admin/main/up/model/datamodel",
        "//app/admin/main/up/util/hbaseutil",
        "//app/interface/main/creative/model/data",
        "//library/cache/memcache",
        "//library/database/hbase.v2:hbase_v2",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/sync/errgroup",
        "//vendor/github.com/tsuna/gohbase/hrpc",
        "//vendor/golang.org/x/net/context",
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
