package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "status",
    srcs = ["status.go"],
    importpath = "go-common/library/net/rpc/warden/status",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/ecode",
        "//library/ecode/pb",
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/github.com/golang/protobuf/ptypes",
        "//vendor/github.com/pkg/errors",
        "//vendor/google.golang.org/grpc/codes",
        "//vendor/google.golang.org/grpc/status",
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

go_test(
    name = "status_test",
    srcs = ["status_test.go"],
    embed = [":status"],
    tags = ["manual"],
    deps = [
        "//library/ecode",
        "//library/ecode/pb",
        "//vendor/github.com/golang/protobuf/ptypes",
        "//vendor/github.com/golang/protobuf/ptypes/timestamp",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/stretchr/testify/assert",
        "//vendor/google.golang.org/grpc/codes",
        "//vendor/google.golang.org/grpc/status",
    ],
)
