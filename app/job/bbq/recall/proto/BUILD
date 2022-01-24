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
    name = "proto_proto",
    srcs = ["forward_index.proto"],
    tags = ["automanaged"],
    deps = [
        "//app/job/bbq/recall/proto/quality:quality_proto",
        "//github.com/gogo/protobuf/gogoproto:gogoproto_proto",
    ],
)

go_proto_library(
    name = "proto_go_proto",
    importpath = "go-common/app/job/bbq/recall/proto",
    proto = ":proto_proto",
    tags = ["automanaged"],
    deps = [
        "//app/job/bbq/recall/proto/quality",
        "//github.com/gogo/protobuf/gogoproto:gogo_proto",
    ],
)

go_library(
    name = "proto",
    embed = [":proto_go_proto"],
    importpath = "go-common/app/job/bbq/recall/proto",
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
    srcs = [
        ":package-srcs",
        "//app/job/bbq/recall/proto/quality:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
