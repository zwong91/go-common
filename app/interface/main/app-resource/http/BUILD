package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "abtest.go",
        "audit.go",
        "broadcast.go",
        "domain.go",
        "http.go",
        "interest.go",
        "local.go",
        "module.go",
        "notice.go",
        "param.go",
        "plugin.go",
        "show.go",
        "sidebar.go",
        "splash.go",
        "static.go",
        "version.go",
        "white.go",
    ],
    importpath = "go-common/app/interface/main/app-resource/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-resource/conf",
        "//app/interface/main/app-resource/model",
        "//app/interface/main/app-resource/model/module",
        "//app/interface/main/app-resource/model/version",
        "//app/interface/main/app-resource/service/abtest",
        "//app/interface/main/app-resource/service/audit",
        "//app/interface/main/app-resource/service/broadcast",
        "//app/interface/main/app-resource/service/domain",
        "//app/interface/main/app-resource/service/guide",
        "//app/interface/main/app-resource/service/module",
        "//app/interface/main/app-resource/service/notice",
        "//app/interface/main/app-resource/service/param",
        "//app/interface/main/app-resource/service/ping",
        "//app/interface/main/app-resource/service/plugin",
        "//app/interface/main/app-resource/service/show",
        "//app/interface/main/app-resource/service/sidebar",
        "//app/interface/main/app-resource/service/splash",
        "//app/interface/main/app-resource/service/static",
        "//app/interface/main/app-resource/service/version",
        "//app/interface/main/app-resource/service/white",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//vendor/github.com/golang/protobuf/proto",
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
