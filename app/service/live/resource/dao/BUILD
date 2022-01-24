package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "confkv.go",
        "dao.go",
        "resource.go",
        "titans.go",
        "tree.go",
        "user_resource.go",
    ],
    importpath = "go-common/app/service/live/resource/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/resource/api/grpc/v1:grpc",
        "//app/service/live/resource/api/http/v1:http",
        "//app/service/live/resource/conf",
        "//app/service/live/resource/lrucache",
        "//app/service/live/resource/model",
        "//library/database/orm",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/siddontang/go-mysql/mysql",
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
