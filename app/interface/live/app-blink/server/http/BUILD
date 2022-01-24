package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = ["http.go"],
    importpath = "go-common/app/interface/live/app-blink/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-blink/api/http/v1:http",
        "//app/interface/live/app-blink/conf",
        "//app/interface/live/app-blink/service",
        "//app/interface/live/app-blink/service/v1:service",
        "//app/service/live/resource/api/grpc/v1:grpc",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/verify",
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
