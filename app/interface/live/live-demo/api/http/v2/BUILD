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
    name = "http",
    srcs = ["service.bm.go"],
    embed = [":v2_go_proto"],
    importpath = "go-common/app/interface/live/live-demo/api/http/v2",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
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
    name = "v2_proto",
    srcs = ["service.proto"],
    tags = ["automanaged"],
    deps = [
        "//github.com/gogo/protobuf/gogoproto:gogoproto_proto",
        "@go_googleapis//google/api:annotations_proto",
    ],
)

go_proto_library(
    name = "v2_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "go-common/app/interface/live/live-demo/api/http/v2",
    proto = ":v2_proto",
    tags = ["automanaged"],
    deps = [
        "//github.com/gogo/protobuf/gogoproto:gogo_proto",
        "@go_googleapis//google/api:annotations_go_proto",
    ],
)
