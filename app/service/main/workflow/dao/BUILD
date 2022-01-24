package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "callback_test.go",
        "dao_test.go",
        "es_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/workflow/conf",
        "//app/service/main/workflow/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "callback.go",
        "dao.go",
        "es.go",
    ],
    importpath = "go-common/app/service/main/workflow/dao",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/workflow/conf",
        "//app/service/main/workflow/model",
        "//library/database/elastic",
        "//library/database/orm",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/jinzhu/gorm",
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
    srcs = [
        ":package-srcs",
        "//app/service/main/workflow/dao/sobot:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
