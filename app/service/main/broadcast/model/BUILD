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
    srcs = ["model.proto"],
    tags = ["automanaged"],
)

go_proto_library(
    name = "model_go_proto",
    importpath = "go-common/app/service/main/broadcast/model",
    proto = ":model_proto",
    tags = ["manual"],
)

go_library(
    name = "model",
    srcs = [
        "auth.go",
        "model.go",
        "model_v1.go",
        "online.go",
        "operation.go",
        "room.go",
        "server.go",
    ],
    embed = [":model_go_proto"],
    importpath = "go-common/app/service/main/broadcast/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/broadcast/libs/bufio",
        "//app/service/main/broadcast/libs/bytes",
        "//app/service/main/broadcast/libs/encoding/binary",
        "//app/service/main/broadcast/libs/websocket",
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
