package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "assist_test",
    srcs = [
        "assist_redis_test.go",
        "assist_test.go",
        "dao_test.go",
        "log_test.go",
        "mc_test.go",
    ],
    embed = [":assist"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/assist/conf",
        "//app/service/main/assist/model/assist",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "assist",
    srcs = [
        "assist.go",
        "assist_redis.go",
        "dao.go",
        "log.go",
        "mc.go",
    ],
    importpath = "go-common/app/service/main/assist/dao/assist",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/assist/conf",
        "//app/service/main/assist/model/assist",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/time",
        "//library/xstr",
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
