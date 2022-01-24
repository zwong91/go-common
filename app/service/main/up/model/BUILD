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
    srcs = ["message.proto"],
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "model_go_proto",
    importpath = "go-common/app/service/main/up/model",
    proto = ":model_proto",
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)

go_library(
    name = "model",
    srcs = [
        "archive.go",
        "monitor.go",
        "msg.go",
        "rpc.go",
        "sign_up.go",
        "up.go",
        "up_card.go",
        "up_group.go",
        "up_info.go",
        "up_special.go",
        "up_switch.go",
    ],
    embed = [":model_go_proto"],
    importpath = "go-common/app/service/main/up/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/time",
        "//vendor/github.com/siddontang/go-mysql/mysql",
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
        "//app/service/main/up/model/data:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
