load("@rules_proto//proto:defs.bzl", "proto_library")
load(
    "@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
)

package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "liverpc",
    srcs = [
        "Bubble.liverpc.go",
        "Color.liverpc.go",
        "DanmuConf.liverpc.go",
        "GrayRule.liverpc.go",
        "Remind.liverpc.go",
        "Tag.liverpc.go",
    ],
    embed = [":v1_go_proto"],
    importpath = "go-common/app/service/live/userext/api/liverpc/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/net/rpc/liverpc",
        "//vendor/github.com/golang/protobuf/proto",
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

proto_library(
    name = "v1_proto",
    srcs = [
        "Bubble.proto",
        "Color.proto",
        "DanmuConf.proto",
        "GrayRule.proto",
        "Remind.proto",
        "Tag.proto",
    ],
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "v1_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "go-common/app/service/live/userext/api/liverpc/v1",
    proto = ":v1_proto",
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)
