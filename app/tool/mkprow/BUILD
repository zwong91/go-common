package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "mkprow",
    embed = [":mkprow_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "mkprow_lib",
    srcs = [
        "main.go",
        "template.go",
    ],
    importpath = "go-common/app/tool/mkprow",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/ghodss/yaml",
        "//vendor/github.com/mohae/deepcopy",
        "//vendor/k8s.io/apimachinery/pkg/util/sets",
        "//vendor/k8s.io/test-infra/prow/config",
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
