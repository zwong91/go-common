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
    name = "quality_proto",
    srcs = ["quality.proto"],
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "quality_go_proto",
    importpath = "go-common/app/job/bbq/recall/proto/quality",
    proto = ":quality_proto",
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)

go_library(
    name = "quality",
    embed = [":quality_go_proto"],
    importpath = "go-common/app/job/bbq/recall/proto/quality",
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
