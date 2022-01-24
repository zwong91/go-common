load("@rules_proto//proto:defs.bzl", "proto_library")
load(
    "@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
)

package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

proto_library(
    name = "liverpc_proto",
    srcs = ["liverpc.proto"],
    tags = ["automanaged"],
)

go_proto_library(
    name = "liverpc_go_proto",
    importpath = "go-common/library/net/rpc/liverpc",
    proto = ":liverpc_proto",
    tags = ["automanaged"],
)

go_test(
    name = "liverpc_test",
    srcs = ["client_conn_test.go"],
    embed = [":liverpc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/conf/env",
        "//library/net/metadata",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "liverpc",
    srcs = [
        "client.go",
        "client_conn.go",
        "option.go",
        "protocol.go",
    ],
    embed = [":liverpc_go_proto"],
    importpath = "go-common/library/net/rpc/liverpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/conf/env",
        "//library/log",
        "//library/naming",
        "//library/naming/discovery",
        "//library/net/metadata",
        "//library/net/trace",
        "//library/stat",
        "//library/time",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/github.com/json-iterator/go",
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
    srcs = [
        ":package-srcs",
        "//library/net/rpc/liverpc/context:all-srcs",
        "//library/net/rpc/liverpc/test:all-srcs",
        "//library/net/rpc/liverpc/testdata:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
