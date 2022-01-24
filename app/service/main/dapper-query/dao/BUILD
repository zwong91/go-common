load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "util.go",
    ],
    importpath = "go-common/app/service/main/dapper-query/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/dapper-query/conf",
        "//app/service/main/dapper-query/model",
        "//library/log",
        "//vendor/github.com/dgryski/go-farm",
        "//vendor/github.com/influxdata/influxdb/client/v2:client",
        "//vendor/github.com/tsuna/gohbase",
        "//vendor/github.com/tsuna/gohbase/filter",
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
    name = "dao_test",
    srcs = ["dao_test.go"],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/dapper-query/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/golang.org/x/sys/unix",
    ],
)
