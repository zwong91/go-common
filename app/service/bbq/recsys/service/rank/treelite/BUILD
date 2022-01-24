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
    name = "treelite",
    srcs = ["predict.go"],
    embed = [":treelite_go_proto"],
    importpath = "go-common/app/service/bbq/recsys/service/rank/treelite",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
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
    name = "treelite_proto",
    srcs = ["tree.proto"],
    tags = ["automanaged"],
)

go_proto_library(
    name = "treelite_go_proto",
    importpath = "go-common/app/service/bbq/recsys/service/rank/treelite",
    proto = ":treelite_proto",
    tags = ["automanaged"],
)
