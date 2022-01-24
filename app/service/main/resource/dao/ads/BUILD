package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "ads_test",
    srcs = [
        "dao_test.go",
        "redis_test.go",
        "vdoad_test.go",
    ],
    embed = [":ads"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/resource/conf",
        "//library/cache/redis",
        "//library/database/sql",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "ads",
    srcs = [
        "dao.go",
        "redis.go",
        "vdoad.go",
    ],
    importpath = "go-common/app/service/main/resource/dao/ads",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/resource/conf",
        "//app/service/main/resource/model",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//vendor/github.com/dgryski/go-farm",
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
