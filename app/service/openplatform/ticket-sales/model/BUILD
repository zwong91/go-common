load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

package(default_visibility = ["//visibility:public"])

go_library(
    name = "model",
    srcs = [
        "cache.go",
        "distrib.go",
        "model.go",
        "order.go",
        "pay.go",
        "promo.go",
        "stock.go",
        "ticket.go",
    ],
    importpath = "go-common/app/service/openplatform/ticket-sales/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/common/openplatform/encoding",
        "//app/service/openplatform/ticket-sales/api/grpc/type",
        "//app/service/openplatform/ticket-sales/api/grpc/v1:grpc",
        "//library/log",
        "//library/time",
        "//vendor/github.com/gogo/protobuf/types",
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
        "//app/service/openplatform/ticket-sales/model/consts:all-srcs",
        "//app/service/openplatform/ticket-sales/model/order_checker/account:all-srcs",
        "//app/service/openplatform/ticket-sales/model/order_checker/item:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
