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
    srcs = [
        "api.proto",
        "archive.proto",
        "reply.proto",
        "request.proto",
        "sign_up.proto",
        "up.proto",
    ],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/api:api_proto",
        "//github.com/gogo/protobuf/gogoproto:gogoproto_proto",
        "@go_googleapis//google/api:annotations_proto",
    ],
)

go_proto_library(
    name = "v1_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "go-common/app/service/main/up/api/v1",
    proto = ":v1_proto",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/api",
        "//github.com/gogo/protobuf/gogoproto:gogo_proto",
        "@go_googleapis//google/api:annotations_go_proto",
    ],
)

go_library(
    name = "api",
    srcs = [
        "api.pb.bm.go",
        "client.go",
        "mock.go",
    ],
    embed = [":v1_go_proto"],
    importpath = "go-common/app/service/main/up/api/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/protoc-gen-bm/jsonpb",
        "//library/net/http/blademaster",
        "//library/net/rpc/warden",
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/golang.org/x/net/context",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
