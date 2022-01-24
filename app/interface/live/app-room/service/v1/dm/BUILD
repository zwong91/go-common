package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dm",
    srcs = [
        "dM.go",
        "data.go",
        "validate.go",
    ],
    importpath = "go-common/app/interface/live/app-room/service/v1/dm",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-room/api/http/v1:http",
        "//app/interface/live/app-room/conf",
        "//app/interface/live/app-room/model",
        "//app/service/live/live-dm/api/grpc/v1:grpc",
        "//app/service/live/live_riskcontrol/api/grpc/v1:grpc",
        "//app/service/live/xcaptcha/api/grpc/v1:grpc",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
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
