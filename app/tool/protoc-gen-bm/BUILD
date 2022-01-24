package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "protoc-gen-bm",
    embed = [":protoc-gen-bm_lib"],
    tags = ["automanaged"],
)

go_library(
    name = "protoc-gen-bm_lib",
    srcs = ["main.go"],
    importpath = "go-common/app/tool/protoc-gen-bm",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/protoc-gen-bm/codegenerator",
        "//app/tool/protoc-gen-bm/genbm",
        "//app/tool/protoc-gen-bm/util",
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
        "//app/tool/protoc-gen-bm/codegenerator:all-srcs",
        "//app/tool/protoc-gen-bm/genbm:all-srcs",
        "//app/tool/protoc-gen-bm/generator:all-srcs",
        "//app/tool/protoc-gen-bm/jsonpb:all-srcs",
        "//app/tool/protoc-gen-bm/util:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
