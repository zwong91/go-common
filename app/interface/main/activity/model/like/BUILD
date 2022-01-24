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
    name = "like",
    srcs = [
        "act.go",
        "action.go",
        "like.go",
        "params.go",
        "rpc.go",
        "subject.go",
        "useractievement.go",
    ],
    embed = [":like_go_proto"],
    importpath = "go-common/app/interface/main/activity/model/like",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/archive/api",
        "//library/time",
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

proto_library(
    name = "like_proto",
    srcs = [
        "achievements.proto",
        "match.proto",
        "missiongroup.proto",
        "subject.proto",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "like_go_proto",
    importpath = "go-common/app/interface/main/activity/model/like",
    proto = ":like_proto",
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)
