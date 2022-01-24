package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
    "go_test",
)

go_binary(
    name = "protoc-gen-bm",
    embed = [":protoc-gen-bm_lib"],
    tags = ["automanaged"],
)

go_test(
    name = "protoc-gen-bm_test",
    srcs = [
        "command_line_test.go",
        "generator_test.go",
    ],
    embed = [":protoc-gen-bm_lib"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/plugin",
    ],
)

go_library(
    name = "protoc-gen-bm_lib",
    srcs = [
        "command_line.go",
        "generator.go",
        "go_naming.go",
        "helper.go",
        "main.go",
    ],
    importpath = "go-common/app/tool/bmproto/protoc-gen-bm",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/bmproto/protoc-gen-bm/extensions/gogoproto",
        "//app/tool/liverpc/protoc-gen-liverpc/gen",
        "//app/tool/liverpc/protoc-gen-liverpc/gen/stringutils",
        "//app/tool/liverpc/protoc-gen-liverpc/gen/typemap",
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/descriptor",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/plugin",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/siddontang/go/ioutil2",
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
    srcs = [
        ":package-srcs",
        "//app/tool/bmproto/protoc-gen-bm/example:all-srcs",
        "//app/tool/bmproto/protoc-gen-bm/extensions/gogoproto:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
