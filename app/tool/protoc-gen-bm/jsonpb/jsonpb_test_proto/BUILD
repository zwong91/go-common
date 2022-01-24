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
    name = "jsonpb_proto",
    srcs = [
        "more_test_objects.proto",
        "test_objects.proto",
    ],
    tags = ["manual"],
    deps = [
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:duration_proto",
        "@com_google_protobuf//:struct_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@com_google_protobuf//:wrappers_proto",
    ],
)

go_proto_library(
    name = "jsonpb_go_proto",
    importpath = "go-common/app/tool/protoc-gen-bm/jsonpb/jsonpb_test_proto",
    proto = ":jsonpb_proto",
    tags = ["manual"],
)

go_library(
    name = "jsonpb_test_proto",
    embed = [":jsonpb_go_proto"],
    importpath = "go-common/app/tool/protoc-gen-bm/jsonpb/jsonpb_test_proto",
    tags = ["manual"],
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
