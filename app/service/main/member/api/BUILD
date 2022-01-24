load("@rules_proto//proto:defs.bzl", "proto_library")
load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)
load(
    "@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
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
        "//app/service/main/member/api/gorpc:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

proto_library(
    name = "api_proto",
    srcs = ["api.proto"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "api_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "go-common/app/service/main/member/api",
    proto = ":api_proto",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)

go_library(
    name = "api",
    srcs = [
        "client.go",
        "copy.go",
    ],
    embed = [":api_go_proto"],
    importpath = "go-common/app/service/main/member/api",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//library/net/rpc/warden",
        "//vendor/google.golang.org/grpc",
    ],
)

go_test(
    name = "api_test",
    srcs = ["client_test.go"],
    embed = [":api"],
    tags = ["automanaged"],
    deps = ["//vendor/github.com/smartystreets/goconvey/convey"],
)
