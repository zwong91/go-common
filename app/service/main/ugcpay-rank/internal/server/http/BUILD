package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "elec.go",
        "http.go",
    ],
    importpath = "go-common/app/service/main/ugcpay-rank/internal/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/ugcpay-rank/api/http",
        "//app/service/main/ugcpay-rank/internal/conf",
        "//app/service/main/ugcpay-rank/internal/model",
        "//app/service/main/ugcpay-rank/internal/service",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/verify",
        "//vendor/github.com/json-iterator/go",
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
