load("@rules_proto//proto:defs.bzl", "proto_library")

package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)
load(
    "@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
)

go_library(
    name = "archive",
    srcs = [
        "addit.go",
        "archive.go",
        "pgc.go",
        "report_result.go",
        "rpc.go",
        "rpc2.go",
    ],
    embed = [":archive_go_proto"],
    importpath = "go-common/app/service/main/archive/model/archive",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/api",
        "//library/time",
    ],
)

proto_library(
    name = "archive_proto",
    srcs = [
        "archive.proto",
        "bvc.proto",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "archive_go_proto",
    importpath = "go-common/app/service/main/archive/model/archive",
    proto = ":archive_proto",
    tags = ["automanaged"],
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
