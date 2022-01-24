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
    name = "v1_proto",
    srcs = ["api.proto"],
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_library(
    name = "api",
    srcs = [
        "copy.go",
        "doc.go",
    ],
    embed = [":v1_go_proto"],
    importpath = "go-common/app/tool/gengo/cmd/deepcopy-gen/examples/api/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//app/tool/gengo/cmd/deepcopy-gen/examples/model"],
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

go_proto_library(
    name = "v1_go_proto",
    importpath = "go-common/app/tool/gengo/cmd/deepcopy-gen/examples/api/v1",
    proto = ":v1_proto",
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)
