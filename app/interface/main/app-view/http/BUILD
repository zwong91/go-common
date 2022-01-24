package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "bnj.go",
        "http.go",
        "local.go",
        "report.go",
        "view.go",
    ],
    importpath = "go-common/app/interface/main/app-view/http",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-view/conf",
        "//app/interface/main/app-view/model",
        "//app/interface/main/app-view/model/view",
        "//app/interface/main/app-view/service/report",
        "//app/interface/main/app-view/service/view",
        "//app/service/main/resource/model",
        "//library/ecode",
        "//library/log",
        "//library/log/anticheat",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/proxy",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/metadata",
        "//library/queue/databus",
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
