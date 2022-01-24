package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "card.go",
        "equip.go",
        "group.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/card/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/card/conf",
        "//app/service/main/card/dao",
        "//app/service/main/card/model",
        "//app/service/main/vip/model",
        "//app/service/main/vip/rpc/client",
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
    name = "service_test",
    srcs = [
        "card_test.go",
        "equip_test.go",
        "group_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/card/conf",
        "//app/service/main/card/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
