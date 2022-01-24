package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "goblin",
    srcs = [
        "hotword.go",
        "label.go",
        "playurl.go",
        "service.go",
        "splash.go",
        "upgrade.go",
    ],
    importpath = "go-common/app/interface/main/tv/service/goblin",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/dao/account",
        "//app/interface/main/tv/dao/archive",
        "//app/interface/main/tv/dao/goblin",
        "//app/interface/main/tv/model",
        "//app/interface/main/tv/model/goblin",
        "//app/service/main/archive/api",
        "//app/service/main/tv/api",
        "//library/ecode",
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

go_test(
    name = "goblin_test",
    srcs = ["service_test.go"],
    embed = [":goblin"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tv/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
