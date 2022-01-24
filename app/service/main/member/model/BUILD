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
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "model_go_proto",
    importpath = "go-common/app/service/main/member/model",
    proto = ":model_proto",
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)

go_library(
    name = "model",
    srcs = [
        "base.go",
        "log.go",
        "member.go",
        "moral.go",
        "official.go",
        "property_review.go",
        "realname.go",
        "rpc.go",
        "search.go",
        "user_flag.go",
    ],
    embed = [":model_go_proto"],
    importpath = "go-common/app/service/main/member/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//library/time",
        "//vendor/github.com/satori/go.uuid:go_uuid",
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
        "//app/service/main/member/model/block:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
