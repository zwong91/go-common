package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "bplus_test",
    srcs = [
        "api_test.go",
        "dao_test.go",
        "databus_test.go",
        "redis_test.go",
    ],
    embed = [":bplus"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/model",
        "//app/interface/main/app-interface/model/space",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "bplus",
    srcs = [
        "api.go",
        "dao.go",
        "databus.go",
        "redis.go",
    ],
    importpath = "go-common/app/interface/main/app-interface/dao/bplus",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/model",
        "//app/interface/main/app-interface/model/bplus",
        "//app/interface/main/app-interface/model/space",
        "//library/cache/redis",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/queue/databus",
        "//library/time",
        "//library/xstr",
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
