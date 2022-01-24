load("@rules_proto//proto:defs.bzl", "proto_library")
load(
    "@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
)
load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

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
        "//library/queue/databus/databusutil:all-srcs",
        "//library/queue/databus/metadata:all-srcs",
        "//library/queue/databus/report:all-srcs",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "databus",
    srcs = ["databus.go"],
    embed = [":databus_go_proto"],
    importpath = "go-common/library/queue/databus",
    tags = ["automanaged"],
    deps = [
        "//library/cache/redis",
        "//library/conf/env",
        "//library/container/pool",
        "//library/log",
        "//library/naming",
        "//library/naming/discovery",
        "//library/net/netutil",
        "//library/net/trace",
        "//library/stat/prom",
        "//library/time",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["databus_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/naming/discovery:go_default_library",
        "//library/queue/databus:go_default_library",
        "//library/time:go_default_library",
    ],
)

proto_library(
    name = "databus_proto",
    srcs = ["databus.proto"],
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "databus_go_proto",
    importpath = "go-common/library/queue/databus",
    proto = ":databus_proto",
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)

go_test(
    name = "databus_test",
    srcs = ["databus_test.go"],
    deps = [
        ":databus",
        "//library/naming/discovery",
        "//library/time",
    ],
)
