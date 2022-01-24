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

go_library(
    name = "model",
    srcs = [
        "achieve.go",
        "addit.go",
        "attr.go",
        "audit.go",
        "following.go",
        "log.go",
        "rpc.go",
        "stat.go",
    ],
    embed = [":model_go_proto"],
    importpath = "go-common/app/service/main/relation/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

proto_library(
    name = "model_proto",
    srcs = [
        "model.proto",
    ],
    tags = ["manual"],
    visibility = ["//visibility:public"],
    deps = [
        "//github.com/gogo/protobuf/gogoproto:gogoproto_proto",
    ],
)

go_proto_library(
    name = "model_go_proto",
    importpath = "go-common/app/service/main/relation/model",
    proto = ":model_proto",
    tags = ["manual"],
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
    srcs = [
        ":package-srcs",
        "//app/service/main/relation/model/i64b:all-srcs",
        "//app/service/main/relation/model/sets:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
