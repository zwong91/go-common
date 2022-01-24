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

proto_library(
    name = "gogoproto_proto",
    srcs = ["gogo.proto"],
    tags = ["automanaged"],
    deps = ["@com_google_protobuf//:descriptor_proto"],
)

go_proto_library(
    name = "gogoproto_go_proto",
    importpath = "go-common/app/tool/bmproto/protoc-gen-bm/extensions/gogoproto",
    proto = ":gogoproto_proto",
    tags = ["automanaged"],
)

go_library(
    name = "gogoproto",
    embed = [":gogoproto_go_proto"],
    importpath = "go-common/app/tool/bmproto/protoc-gen-bm/extensions/gogoproto",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
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
