load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "distrib.go",
        "order.go",
        "order_detail.go",
        "pay.go",
        "promotion.go",
        "promotion_group.go",
        "promotion_order.go",
        "service.go",
        "stock.go",
        "ticket.go",
    ],
    importpath = "go-common/app/service/openplatform/ticket-sales/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/common/openplatform/random",
        "//app/service/openplatform/ticket-sales/api/grpc/type",
        "//app/service/openplatform/ticket-sales/api/grpc/v1:grpc",
        "//app/service/openplatform/ticket-sales/conf",
        "//app/service/openplatform/ticket-sales/dao",
        "//app/service/openplatform/ticket-sales/model",
        "//app/service/openplatform/ticket-sales/model/consts",
        "//app/service/openplatform/ticket-sales/model/order_checker/account",
        "//app/service/openplatform/ticket-sales/model/order_checker/item",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//vendor/github.com/pkg/errors",
    ],
)

go_test(
    name = "service_test",
    srcs = [
        "distrib_test.go",
        "order_test.go",
        "promotion_group_test.go",
        "promotion_order_test.go",
        "promotion_test.go",
        "service_test.go",
        "stock_test.go",
        "ticket_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/openplatform/ticket-sales/api/grpc/type",
        "//app/service/openplatform/ticket-sales/api/grpc/v1:grpc",
        "//app/service/openplatform/ticket-sales/conf",
        "//app/service/openplatform/ticket-sales/model",
        "//app/service/openplatform/ticket-sales/model/consts",
        "//library/conf/paladin",
        "//library/ecode",
        "//vendor/github.com/davecgh/go-spew/spew",
        "//vendor/github.com/smartystreets/goconvey/convey",
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
        "//app/service/openplatform/ticket-sales/service/mis:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
