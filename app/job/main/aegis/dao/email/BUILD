package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "email",
    srcs = ["dao.go"],
    importpath = "go-common/app/job/main/aegis/dao/email",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/aegis/conf",
        "//app/job/main/aegis/model",
        "//library/cache/redis",
        "//library/log",
        "//vendor/gopkg.in/gomail.v2:gomail_v2",
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
    name = "email_test",
    srcs = ["dao_test.go"],
    embed = [":email"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/aegis/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
