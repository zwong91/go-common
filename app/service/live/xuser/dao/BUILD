package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = ["api_test.go"],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/xuser/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = ["api.go"],
    importpath = "go-common/app/service/live/xuser/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/banned_service/api/liverpc",
        "//app/service/live/fans_medal/api/liverpc",
        "//app/service/live/room/api/liverpc",
        "//app/service/live/xuser/conf",
        "//library/net/rpc/liverpc",
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
        "//app/service/live/xuser/dao/account:all-srcs",
        "//app/service/live/xuser/dao/exp:all-srcs",
        "//app/service/live/xuser/dao/guard:all-srcs",
        "//app/service/live/xuser/dao/roomAdmin:all-srcs",
        "//app/service/live/xuser/dao/vip:all-srcs",
        "//app/service/live/xuser/dao/xanchor:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
