package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "abtest_test.go",
        "cache_test.go",
        "group_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/openplatform/abtest/conf",
        "//app/service/openplatform/abtest/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "abtest.go",
        "cache.go",
        "group.go",
        "service.go",
    ],
    importpath = "go-common/app/service/openplatform/abtest/service",
    tags = ["automanaged"],
    deps = [
        "//app/service/openplatform/abtest/conf",
        "//app/service/openplatform/abtest/dao",
        "//app/service/openplatform/abtest/model",
        "//app/service/openplatform/abtest/model/jump",
        "//library/cache/redis",
        "//library/log",
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
