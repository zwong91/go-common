package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "grpc_test",
    srcs = ["server_test.go"],
    embed = [":grpc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/openplatform/ticket-item/api/grpc/v1:grpc",
        "//app/service/openplatform/ticket-item/conf",
        "//app/service/openplatform/ticket-item/model",
        "//app/service/openplatform/ticket-item/service",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "grpc",
    srcs = ["server.go"],
    importpath = "go-common/app/service/openplatform/ticket-item/server/grpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/openplatform/ticket-item/api/grpc/v1:grpc",
        "//app/service/openplatform/ticket-item/conf",
        "//app/service/openplatform/ticket-item/service",
        "//library/log",
        "//library/net/rpc/warden",
        "//vendor/google.golang.org/grpc",
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
