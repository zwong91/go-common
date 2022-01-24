package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "search_test",
    srcs = [
        "dao_test.go",
        "log_test.go",
        "record_test.go",
    ],
    embed = [":search"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/reply/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "search",
    srcs = [
        "dao.go",
        "log.go",
        "record.go",
    ],
    importpath = "go-common/app/interface/main/reply/dao/search",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/reply/conf",
        "//app/interface/main/reply/model/adminlog",
        "//app/interface/main/reply/model/reply",
        "//library/database/elastic",
        "//library/log",
        "//library/net/http/blademaster",
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
