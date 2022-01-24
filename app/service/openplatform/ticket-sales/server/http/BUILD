load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "http",
    srcs = [
        "distrib.go",
        "http.go",
        "promotion.go",
        "promotion_group.go",
        "promotion_order.go",
        "settle.go",
    ],
    importpath = "go-common/app/service/openplatform/ticket-sales/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/openplatform/ticket-sales/api/grpc/v1:grpc",
        "//app/service/openplatform/ticket-sales/conf",
        "//app/service/openplatform/ticket-sales/model",
        "//app/service/openplatform/ticket-sales/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/render",
        "//vendor/github.com/pkg/errors",
    ],
)

go_test(
    name = "http_test",
    srcs = ["http_test.go"],
    embed = [":http"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/openplatform/ticket-sales/conf",
        "//app/service/openplatform/ticket-sales/service",
        "//library/conf/paladin",
        "//library/net/http/blademaster",
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
