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
    name = "model_proto",
    srcs = ["dm.proto"],
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "model_go_proto",
    importpath = "go-common/app/job/main/dm/model",
    proto = ":model_proto",
    tags = ["manual"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)

go_library(
    name = "model",
    srcs = [
        "dm.go",
        "job.go",
        "subject.go",
    ],
    embed = [":model_go_proto"],
    importpath = "go-common/app/job/main/dm/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/time"],
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
