package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
    "go_test",
)

go_binary(
    name = "protoc-gen-liverpc",
    embed = [":protoc-gen-liverpc_lib"],
    tags = ["automanaged"],
)

go_test(
    name = "protoc-gen-liverpc_test",
    srcs = [
        "command_line_test.go",
        "generator_test.go",
    ],
    embed = [":protoc-gen-liverpc_lib"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/plugin",
    ],
)

go_library(
    name = "protoc-gen-liverpc_lib",
    srcs = [
        "command_line.go",
        "generator.go",
        "go_naming.go",
        "main.go",
    ],
    importpath = "go-common/app/tool/liverpc/protoc-gen-liverpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/liverpc/protoc-gen-liverpc/gen",
        "//app/tool/liverpc/protoc-gen-liverpc/gen/stringutils",
        "//app/tool/liverpc/protoc-gen-liverpc/gen/typemap",
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/descriptor",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/plugin",
        "//vendor/github.com/pkg/errors",
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
        "//app/tool/liverpc/protoc-gen-liverpc/gen:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
