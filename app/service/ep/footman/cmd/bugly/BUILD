package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "bugly",
    embed = [":bugly_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "bugly_lib",
    srcs = ["bugly.go"],
    importpath = "go-common/app/service/ep/footman/cmd/bugly",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/ep/footman/conf",
        "//app/service/ep/footman/service",
        "//library/cache/memcache",
        "//library/container/pool",
        "//library/database/orm",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/netutil/breaker",
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
