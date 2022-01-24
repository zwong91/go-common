package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "mail-api-client",
    srcs = ["mail_api_client.go"],
    importpath = "go-common/app/service/main/vip/dao/mail-api-client",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/net/http/blademaster",
        "//vendor/github.com/pkg/errors",
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
    name = "mail-api-client_test",
    srcs = ["mail_api_client_test.go"],
    embed = [":mail-api-client"],
    tags = ["automanaged"],
    deps = [
        "//library/net/http/blademaster",
        "//library/net/netutil/breaker",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
