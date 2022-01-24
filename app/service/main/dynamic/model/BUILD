load("@rules_proto//proto:defs.bzl", "proto_library")

package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)
load(
    "@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
)

go_library(
    name = "model",
    srcs = [
        "bigdata.go",
        "dynamic.go",
        "rpc.go",
    ],
    embed = [":model_go_proto"],
    importpath = "go-common/app/service/main/dynamic/model",
    tags = ["automanaged"],
    deps = ["//app/service/main/archive/api"],
)

proto_library(
    name = "model_proto",
    srcs = ["dynamic.proto"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "model_go_proto",
    importpath = "go-common/app/service/main/dynamic/model",
    proto = ":model_proto",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
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
