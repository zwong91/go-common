package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "middleware_test",
    srcs = [
        "config_test.go",
        "response_test.go",
    ],
    embed = [":middleware"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/aegis/model",
        "//app/admin/main/aegis/model/resource",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "middleware",
    srcs = [
        "config.go",
        "response.go",
    ],
    importpath = "go-common/app/admin/main/aegis/model/middleware",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/log"],
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
