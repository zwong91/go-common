package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "aids.go",
        "audit.go",
        "bfs.go",
        "bottom.go",
        "http.go",
        "language.go",
        "local.go",
        "notice.go",
        "wall.go",
    ],
    importpath = "go-common/app/admin/main/app/http",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/app/conf",
        "//app/admin/main/app/model",
        "//app/admin/main/app/model/aids",
        "//app/admin/main/app/model/audit",
        "//app/admin/main/app/model/bottom",
        "//app/admin/main/app/model/language",
        "//app/admin/main/app/model/notice",
        "//app/admin/main/app/model/wall",
        "//app/admin/main/app/service/aids",
        "//app/admin/main/app/service/audit",
        "//app/admin/main/app/service/bfs",
        "//app/admin/main/app/service/bottom",
        "//app/admin/main/app/service/language",
        "//app/admin/main/app/service/notice",
        "//app/admin/main/app/service/ping",
        "//app/admin/main/app/service/wall",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/http/blademaster/render",
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
