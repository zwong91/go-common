load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    data = ["msm-service-example.toml"],
    importpath = "go-common/app/service/main/msm/cmd",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/msm/conf",
        "//app/service/main/msm/http",
        "//app/service/main/msm/service",
        "//library/log",
    ],
)

go_binary(
    name = "cmd",
    embed = [":cmd_lib"],
    tags = ["automanaged"],
)
