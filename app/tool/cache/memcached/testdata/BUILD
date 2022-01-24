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
    name = "model_proto",
    srcs = ["model.proto"],
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "model_go_proto",
    importpath = "go-common/app/tool/cache/memcached/testdata",
    proto = ":model_proto",
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)

go_test(
    name = "testdata_test",
    srcs = ["dao_test.go"],
    embed = [":testdata"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "testdata",
    srcs = [
        "dao.go",
        "mc.cache.go",
    ],
    embed = [":model_go_proto"],
    importpath = "go-common/app/tool/cache/memcached/testdata",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/cache/memcache",
        "//library/container/pool",
        "//library/log",
        "//library/stat/prom",
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
