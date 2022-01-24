package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = ["dao_test.go"],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/block/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "http.go",
        "mc.go",
        "mysql.go",
        "rpc.go",
    ],
    importpath = "go-common/app/admin/main/block/dao",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/block/conf",
        "//app/admin/main/block/model",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//app/service/main/figure/model",
        "//app/service/main/figure/rpc/client",
        "//app/service/main/spy/model",
        "//app/service/main/spy/rpc/client",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
        "//library/net/http/blademaster",
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
