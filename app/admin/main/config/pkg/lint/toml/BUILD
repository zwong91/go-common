package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "toml_test",
    srcs = [
        "testdata_test.go",
        "tomllint_test.go",
    ],
    embed = [":toml"],
    tags = ["automanaged"],
)

go_library(
    name = "toml",
    srcs = ["tomllint.go"],
    importpath = "go-common/app/admin/main/config/pkg/lint/toml",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/config/pkg/lint",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
