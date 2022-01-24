load("@rules_proto//proto:defs.bzl", "proto_library")
load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)
load(
    "@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
)

proto_library(
    name = "v1_proto",
    srcs = ["api.proto"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "v1_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "go-common/app/service/main/identify/api/grpc",
    proto = ":v1_proto",
    tags = ["manual"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)

go_library(
    name = "grpc",
    srcs = [
        "client.go",
        "string.go",
    ],
    embed = [":v1_go_proto"],
    importpath = "go-common/app/service/main/identify/api/grpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/net/rpc/warden",
        "//vendor/github.com/gogo/protobuf/proto",
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

go_test(
    name = "grpc_test",
    srcs = ["string_test.go"],
    embed = [":grpc"],
    rundir = ".",
    tags = ["automanaged"],
)
