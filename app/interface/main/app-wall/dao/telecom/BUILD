package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "telecom_test",
    srcs = ["telecom_test.go"],
    embed = [":telecom"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-wall/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "telecom",
    srcs = [
        "api.go",
        "cache.go",
        "memcache.go",
        "sms.go",
        "telecom.go",
    ],
    importpath = "go-common/app/interface/main/app-wall/dao/telecom",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-wall/conf",
        "//app/interface/main/app-wall/model/telecom",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/xxtea/xxtea-go/xxtea",
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
