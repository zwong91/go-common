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
    name = "types_proto",
    srcs = ["status.proto"],
    tags = ["automanaged"],
    deps = ["@com_google_protobuf//:any_proto"],
)

go_proto_library(
    name = "types_go_proto",
    importpath = "go-common/library/ecode/internal/types",
    proto = ":types_proto",
    tags = ["automanaged"],
)

go_library(
    name = "types",
    embed = [":types_go_proto"],
    importpath = "go-common/library/ecode/internal/types",
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
