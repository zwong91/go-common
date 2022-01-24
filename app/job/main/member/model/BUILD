load("@rules_proto//proto:defs.bzl", "proto_library")

package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)
load(
    "@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
)

proto_library(
    name = "model_proto",
    srcs = ["member.proto"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "model_go_proto",
    importpath = "go-common/app/job/main/member/model",
    proto = ":model_proto",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)

go_test(
    name = "model_test",
    srcs = ["model_test.go"],
    embed = [":model"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/time",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "model",
    srcs = [
        "base.go",
        "model.go",
        "moral.go",
        "notify.go",
        "official.go",
        "old_account.go",
        "realname.go",
    ],
    embed = [":model_go_proto"],
    importpath = "go-common/app/job/main/member/model",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/member/conf",
        "//library/log",
        "//library/time",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/satori/go.uuid:go_uuid",
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
        "//app/job/main/member/model/block:all-srcs",
        "//app/job/main/member/model/queue:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
