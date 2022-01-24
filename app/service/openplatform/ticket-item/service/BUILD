package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "guest_test.go",
        "item_test.go",
        "seat_test.go",
        "service_test.go",
        "venue_test.go",
        "version_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/openplatform/ticket-item/api/grpc/v1:grpc",
        "//app/service/openplatform/ticket-item/conf",
        "//app/service/openplatform/ticket-item/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "area.go",
        "bulletin.go",
        "guest.go",
        "item.go",
        "place.go",
        "seat.go",
        "service.go",
        "venue.go",
        "version.go",
    ],
    importpath = "go-common/app/service/openplatform/ticket-item/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/openplatform/ticket-item/api/grpc/v1:grpc",
        "//app/service/openplatform/ticket-item/conf",
        "//app/service/openplatform/ticket-item/dao",
        "//app/service/openplatform/ticket-item/model",
        "//library/ecode",
        "//library/log",
        "//library/time",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/gopkg.in/go-playground/validator.v9:validator_v9",
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
