package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "deepcopy-gen",
    embed = [":deepcopy-gen_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "deepcopy-gen_lib",
    srcs = ["main.go"],
    importpath = "go-common/app/tool/gengo/cmd/deepcopy-gen",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/gengo/args",
        "//app/tool/gengo/cmd/deepcopy-gen/generators",
        "//vendor/github.com/golang/glog",
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
        "//app/tool/gengo/cmd/deepcopy-gen/examples/api/v1:all-srcs",
        "//app/tool/gengo/cmd/deepcopy-gen/examples/model:all-srcs",
        "//app/tool/gengo/cmd/deepcopy-gen/examples/tests:all-srcs",
        "//app/tool/gengo/cmd/deepcopy-gen/generators:all-srcs",
        "//app/tool/gengo/cmd/deepcopy-gen/sets:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
