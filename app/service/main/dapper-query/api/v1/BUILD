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

go_library(
    name = "api",
    srcs = [
        "api.pb.bm.go",
        "marshal.go",
    ],
    embed = [":v1_go_proto"],
    importpath = "go-common/app/service/main/dapper-query/api/v1",
    tags = ["manual"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/protoc-gen-bm/jsonpb",
        "//library/net/http/blademaster",
    ],
)

proto_library(
    name = "v1_proto",
    srcs = ["api.proto"],
    deps = [
        "//github.com/gogo/protobuf/gogoproto:gogoproto_proto",
        "@go_googleapis//google/api:annotations_proto",
    ],
)

go_proto_library(
    name = "v1_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "go-common/app/service/main/dapper-query/api/v1",
    proto = ":v1_proto",
    deps = [
        "//github.com/gogo/protobuf/gogoproto:gogo_proto",
        "@go_googleapis//google/api:annotations_go_proto",
    ],
)
