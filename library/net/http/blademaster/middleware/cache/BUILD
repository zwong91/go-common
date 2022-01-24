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
    name = "cache_proto",
    srcs = ["page.proto"],
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "cache_go_proto",
    importpath = "go-common/library/net/http/blademaster/middleware/cache",
    proto = ":cache_proto",
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)

go_test(
    name = "cache_test",
    srcs = [
        "cache_test.go",
        "example_test.go",
    ],
    embed = [":cache"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/cache/memcache",
        "//library/container/pool",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/cache/store",
        "//library/time",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "cache",
    srcs = [
        "cache.go",
        "control.go",
        "degrade.go",
        "page.go",
    ],
    embed = [":cache_go_proto"],
    importpath = "go-common/library/net/http/blademaster/middleware/cache",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/cache/store",
        "//vendor/github.com/gogo/protobuf/proto",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/cache/memcache:go_default_library",
        "//library/container/pool:go_default_library",
        "//library/ecode:go_default_library",
        "//library/net/http/blademaster:go_default_library",
        "//library/net/http/blademaster/middleware/cache:go_default_library",
        "//library/net/http/blademaster/middleware/cache/store:go_default_library",
        "//library/time:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
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
        "//library/net/http/blademaster/middleware/cache/store:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
