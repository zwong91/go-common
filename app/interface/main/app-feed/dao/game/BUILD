package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "game_test",
    srcs = ["dao_test.go"],
    embed = [":game"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-card/model/card/operate",
        "//app/interface/main/app-feed/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "game",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/main/app-feed/dao/game",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-card/model/card/operate",
        "//app/interface/main/app-feed/conf",
        "//library/database/sql",
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
