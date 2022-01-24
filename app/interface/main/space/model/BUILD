load("@rules_proto//proto:defs.bzl", "proto_library")
load(
    "@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
)

package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "model",
    srcs = [
        "account.go",
        "album.go",
        "app.go",
        "archive.go",
        "article.go",
        "audio.go",
        "bangumi.go",
        "channel.go",
        "coin.go",
        "dynamic.go",
        "elec.go",
        "fav.go",
        "game.go",
        "live.go",
        "setting.go",
        "shop.go",
        "tag.go",
        "theme.go",
        "web.go",
    ],
    embed = [":model_go_proto"],
    importpath = "go-common/app/interface/main/space/model",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/favorite/model",
        "//library/time",
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

proto_library(
    name = "model_proto",
    srcs = ["proto.proto"],
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "model_go_proto",
    importpath = "go-common/app/interface/main/space/model",
    proto = ":model_proto",
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)
