package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "abtest_test.go",
        "dao_test.go",
        "group_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/openplatform/abtest/conf",
        "//app/service/openplatform/abtest/model",
        "//library/cache/redis",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "abtest.go",
        "dao.go",
        "group.go",
        "redis.go",
    ],
    importpath = "go-common/app/service/openplatform/abtest/dao",
    tags = ["automanaged"],
    deps = [
        "//app/service/openplatform/abtest/conf",
        "//app/service/openplatform/abtest/model",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
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
