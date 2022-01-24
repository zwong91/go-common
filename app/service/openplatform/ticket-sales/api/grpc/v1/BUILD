load("@rules_proto//proto:defs.bzl", "proto_library")
load(
    "@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
)
load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "grpc",
    srcs = ["client.go"],
    embed = [":v1_go_proto"],
    importpath = "go-common/app/service/openplatform/ticket-sales/api/grpc/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/net/rpc/warden",
        "//vendor/google.golang.org/grpc",
    ],
)

proto_library(
    name = "v1_proto",
    srcs = [
        "order.proto",
        "pay.proto",
        "promotion.proto",
        "promotion_mis.proto",
        "ticket.proto",
    ],
    tags = ["manual"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/openplatform/ticket-sales/api/grpc/type:_type_proto",
        "//github.com/gogo/protobuf/gogoproto:gogoproto_proto",
    ],
)

go_proto_library(
    name = "v1_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "go-common/app/service/openplatform/ticket-sales/api/grpc/v1",
    proto = ":v1_proto",
    tags = ["manual"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/openplatform/ticket-sales/api/grpc/type",
        "//github.com/gogo/protobuf/gogoproto:gogo_proto",
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
