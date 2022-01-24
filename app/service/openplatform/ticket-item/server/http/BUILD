package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "area.go",
        "bulletin.go",
        "guest.go",
        "http.go",
        "item.go",
        "place.go",
        "seat.go",
        "venue.go",
        "version.go",
    ],
    importpath = "go-common/app/service/openplatform/ticket-item/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/openplatform/ticket-item/api/grpc/v1:grpc",
        "//app/service/openplatform/ticket-item/conf",
        "//app/service/openplatform/ticket-item/model",
        "//app/service/openplatform/ticket-item/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//vendor/github.com/pkg/errors",
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
