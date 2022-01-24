package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "card",
    srcs = [
        "card.go",
        "dao.go",
        "rpc.go",
        "up_base.go",
    ],
    importpath = "go-common/app/service/main/up/dao/card",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/up/conf",
        "//app/service/main/up/dao",
        "//app/service/main/up/dao/global",
        "//app/service/main/up/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup",
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
    name = "card_test",
    srcs = [
        "card_test.go",
        "dao_test.go",
        "rpc_test.go",
        "up_base_test.go",
    ],
    embed = [":card"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/up/conf",
        "//app/service/main/up/dao",
        "//app/service/main/up/dao/global",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
