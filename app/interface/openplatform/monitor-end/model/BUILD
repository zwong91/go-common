package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "model",
    srcs = ["params.go"],
    importpath = "go-common/app/interface/openplatform/monitor-end/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/time"],
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
        "//app/interface/openplatform/monitor-end/model/kafka:all-srcs",
        "//app/interface/openplatform/monitor-end/model/monitor:all-srcs",
        "//app/interface/openplatform/monitor-end/model/prom:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
