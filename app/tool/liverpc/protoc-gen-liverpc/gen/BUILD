package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "gen",
    srcs = [
        "logging.go",
        "main.go",
        "version.go",
        "wrappers.go",
    ],
    importpath = "go-common/app/tool/liverpc/protoc-gen-liverpc/gen",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/liverpc/protoc-gen-liverpc/gen/stringutils",
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/descriptor",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/plugin",
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
        "//app/tool/liverpc/protoc-gen-liverpc/gen/stringutils:all-srcs",
        "//app/tool/liverpc/protoc-gen-liverpc/gen/typemap:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
