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
    name = "timemachine_proto",
    srcs = ["timemachine.proto"],
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "timemachine_go_proto",
    importpath = "go-common/app/interface/main/activity/model/timemachine",
    proto = ":timemachine_proto",
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)

go_library(
    name = "timemachine",
    srcs = ["model.go"],
    embed = [":timemachine_go_proto"],
    importpath = "go-common/app/interface/main/activity/model/timemachine",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//app/service/main/archive/api"],
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
