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

go_test(
    name = "model_test",
    srcs = ["article_test.go"],
    embed = [":model"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/smartystreets/goconvey/convey"],
)

go_library(
    name = "model",
    srcs = [
        "activity.go",
        "apply.go",
        "article.go",
        "author.go",
        "cards.go",
        "creation.go",
        "hotspots.go",
        "infoc.go",
        "list.go",
        "rank.go",
        "rpc.go",
        "setting.go",
        "sort.go",
    ],
    embed = [":model_go_proto"],
    importpath = "go-common/app/interface/openplatform/article/model",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/model/data",
        "//app/service/main/account/model",
        "//library/time",
    ],
)

proto_library(
    name = "model_proto",
    srcs = ["model.proto"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "model_go_proto",
    importpath = "go-common/app/interface/openplatform/article/model",
    proto = ":model_proto",
    tags = ["manual"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
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
        "//app/interface/openplatform/article/model/search:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
