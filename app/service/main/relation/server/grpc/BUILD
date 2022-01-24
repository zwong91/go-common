load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "grpc",
    srcs = ["server.go"],
    importpath = "go-common/app/service/main/relation/server/grpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/relation/api",
        "//app/service/main/relation/conf",
        "//app/service/main/relation/model",
        "//app/service/main/relation/service",
        "//library/net/rpc/warden",
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
    name = "grpc_test",
    srcs = ["server_test.go"],
    embed = [":grpc"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/relation/api",
        "//app/service/main/relation/conf",
        "//app/service/main/relation/model",
        "//app/service/main/relation/service",
        "//library/net/rpc/warden",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
