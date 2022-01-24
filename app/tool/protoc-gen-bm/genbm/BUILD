package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "genbm",
    srcs = [
        "bm_generate.go",
        "genbm.go",
        "http_descriptor.go",
    ],
    importpath = "go-common/app/tool/protoc-gen-bm/genbm",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/protoc-gen-bm/generator",
        "//app/tool/protoc-gen-bm/util",
        "//vendor/github.com/golang/glog",
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/descriptor",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/plugin",
        "//vendor/google.golang.org/genproto/googleapis/api/annotations",
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
