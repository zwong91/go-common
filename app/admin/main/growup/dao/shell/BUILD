package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "shell",
    srcs = [
        "shellrequest.go",
        "sign.go",
    ],
    importpath = "go-common/app/admin/main/growup/dao/shell",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/growup/conf",
        "//library/log",
        "//library/net/http/blademaster",
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
    name = "shell_test",
    srcs = [
        "shellrequest_test.go",
        "sign_test.go",
    ],
    embed = [":shell"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/growup/conf",
        "//library/net/http/blademaster",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
