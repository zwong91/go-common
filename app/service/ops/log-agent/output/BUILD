package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "output",
    srcs = [
        "conf.go",
        "output.go",
    ],
    importpath = "go-common/app/service/ops/log-agent/output",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/ops/log-agent/conf/configcenter",
        "//app/service/ops/log-agent/event",
        "//library/log",
        "//vendor/github.com/BurntSushi/toml",
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
        "//app/service/ops/log-agent/output/cache:all-srcs",
        "//app/service/ops/log-agent/output/lancergrpc:all-srcs",
        "//app/service/ops/log-agent/output/lancerlogstream:all-srcs",
        "//app/service/ops/log-agent/output/stdout:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
