load("@rules_proto//proto:defs.bzl", "proto_library")
load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)
load(
    "@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
)

proto_library(
    name = "render_proto",
    srcs = ["pb.proto"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//github.com/gogo/protobuf/gogoproto:gogoproto_proto",
        "@com_google_protobuf//:any_proto",
    ],
)

go_proto_library(
    name = "render_go_proto",
    importpath = "go-common/library/net/http/blademaster/render",
    proto = ":render_proto",
    tags = ["manual"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)

go_library(
    name = "render",
    srcs = [
        "data.go",
        "json.go",
        "protobuf.go",
        "redirect.go",
        "render.go",
        "string.go",
        "xml.go",
    ],
    embed = [":render_go_proto"],
    importpath = "go-common/library/net/http/blademaster/render",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/pkg/errors",
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
