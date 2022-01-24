package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "bulletin_test.go",
        "dao_test.go",
        "guest_test.go",
        "item_test.go",
        "promotion_test.go",
        "screen_test.go",
        "seat_test.go",
        "stock_test.go",
        "tag_test.go",
        "ticket_test.go",
        "venue_test.go",
        "version_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/openplatform/ticket-item/api/grpc/v1:grpc",
        "//app/service/openplatform/ticket-item/conf",
        "//app/service/openplatform/ticket-item/model",
        "//library/log",
        "//vendor/github.com/davecgh/go-spew/spew",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "banner.go",
        "bulletin.go",
        "dao.cache.go",
        "dao.go",
        "fav.go",
        "guest.go",
        "item.go",
        "project.go",
        "promotion.go",
        "screen.go",
        "seat.go",
        "stock.go",
        "tag.go",
        "ticket.go",
        "venue.go",
        "version.go",
        "wish.go",
    ],
    importpath = "go-common/app/service/openplatform/ticket-item/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/common/openplatform/random",
        "//app/service/openplatform/ticket-item/api/grpc/v1:grpc",
        "//app/service/openplatform/ticket-item/conf",
        "//app/service/openplatform/ticket-item/model",
        "//library/cache/redis",
        "//library/database/elastic",
        "//library/database/orm",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
        "//library/time",
        "//vendor/github.com/jinzhu/gorm",
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
