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
    name = "api_proto",
    srcs = ["api.proto"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/favorite/model:model_proto",
        "//github.com/gogo/protobuf/gogoproto:gogoproto_proto",
        "@com_google_protobuf//:empty_proto",
    ],
)

go_proto_library(
    name = "api_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "go-common/app/service/main/favorite/api",
    proto = ":api_proto",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/favorite/model",
        "//github.com/gogo/protobuf/gogoproto:gogo_proto",
    ],
)

go_library(
    name = "api",
    srcs = ["client.go"],
    embed = [":api_go_proto"],
    importpath = "go-common/app/service/main/favorite/api",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/net/rpc/warden",
        "//vendor/google.golang.org/grpc",
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
    srcs = [
        ":package-srcs",
        "//app/service/main/favorite/api/gorpc:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
