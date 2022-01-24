package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "music",
    srcs = ["dao.go"],
    importpath = "go-common/app/job/main/favorite/dao/music",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/favorite/conf",
        "//app/job/main/favorite/model",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
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

go_test(
    name = "music_test",
    srcs = ["dao_test.go"],
    embed = [":music"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/favorite/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
