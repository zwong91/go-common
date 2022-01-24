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
        "associate_banner.go",
        "associate_bind.go",
        "associate_ele_api.go",
        "associate_model.go",
        "associate_old.go",
        "associate_open.go",
        "associate_order.go",
        "associate_panel.go",
        "associate_params.go",
        "bp.go",
        "coupon.go",
        "dialog.go",
        "http.go",
        "jointly.go",
        "mail_api.go",
        "model.go",
        "notify.go",
        "oldvip.go",
        "order.go",
        "panel.go",
        "params.go",
        "pay.go",
        "platform.go",
        "point.go",
        "privilege.go",
        "resource.go",
        "rpc.go",
        "tips.go",
        "user.go",
        "vip.go",
        "welfare.go",
    ],
    embed = [":model_go_proto"],
    importpath = "go-common/app/service/main/vip/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/vip/model",
        "//app/service/main/coupon/api",
        "//app/service/main/coupon/model",
        "//library/time",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

proto_library(
    name = "model_proto",
    srcs = ["vip.proto"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "model_go_proto",
    importpath = "go-common/app/service/main/vip/model",
    proto = ":model_proto",
    tags = ["automanaged"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogo_proto"],
)
