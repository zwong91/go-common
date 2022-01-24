package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "http.go",
        "ip.go",
        "local.go",
        "mobile.go",
        "offer.go",
        "operator.go",
        "telecom.go",
        "unicom.go",
        "wall.go",
    ],
    importpath = "go-common/app/interface/main/app-wall/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-wall/conf",
        "//app/interface/main/app-wall/dao/padding",
        "//app/interface/main/app-wall/model",
        "//app/interface/main/app-wall/model/mobile",
        "//app/interface/main/app-wall/model/telecom",
        "//app/interface/main/app-wall/model/unicom",
        "//app/interface/main/app-wall/service/mobile",
        "//app/interface/main/app-wall/service/offer",
        "//app/interface/main/app-wall/service/operator",
        "//app/interface/main/app-wall/service/ping",
        "//app/interface/main/app-wall/service/telecom",
        "//app/interface/main/app-wall/service/unicom",
        "//app/interface/main/app-wall/service/wall",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/proxy",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/http/blademaster/render",
        "//library/net/metadata",
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
