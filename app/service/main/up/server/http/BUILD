package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "http.go",
        "up.go",
        "up_card.go",
        "up_group.go",
        "up_stat.go",
        "up_switch.go",
    ],
    importpath = "go-common/app/service/main/up/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/up/api/v1:api",
        "//app/service/main/up/conf",
        "//app/service/main/up/model",
        "//app/service/main/up/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/metadata",
        "//vendor/github.com/go-sql-driver/mysql",
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
