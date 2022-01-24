package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "web",
    srcs = [
        "dao.go",
        "dynamic.go",
        "ugc.go",
    ],
    importpath = "go-common/app/job/main/web-goblin/dao/web",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/web-goblin/conf",
        "//library/database/elastic",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
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
    name = "web_test",
    srcs = [
        "dao_test.go",
        "dynamic_test.go",
        "ugc_test.go",
    ],
    embed = [":web"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/web-goblin/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
