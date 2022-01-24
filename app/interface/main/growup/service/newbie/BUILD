package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "newbie",
    srcs = [
        "letter.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/growup/service/newbie",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/growup/conf",
        "//app/interface/main/growup/dao/newbiedao",
        "//app/interface/main/growup/model",
        "//app/service/main/account/api",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup.v2:errgroup_v2",
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
    name = "newbie_test",
    srcs = [
        "letter_test.go",
        "service_test.go",
    ],
    embed = [":newbie"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/growup/conf",
        "//app/interface/main/growup/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
