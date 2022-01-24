package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["client_test.go"],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/infra/config/conf",
        "//app/infra/config/model",
        "//vendor/github.com/BurntSushi/toml",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "cache_tag.go",
        "client.go",
        "service.go",
    ],
    importpath = "go-common/app/infra/config/service/v2",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/infra/config/conf",
        "//app/infra/config/dao/v2:dao",
        "//app/infra/config/model",
        "//library/ecode",
        "//library/log",
        "//library/time",
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
