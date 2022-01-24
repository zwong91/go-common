package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "channel.go",
        "full.go",
        "http.go",
        "pgc.go",
        "recruit.go",
        "share.go",
        "ugc.go",
        "wechat.go",
    ],
    importpath = "go-common/app/interface/main/web-goblin/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/web-goblin/conf",
        "//app/interface/main/web-goblin/model/web",
        "//app/interface/main/web-goblin/service/share",
        "//app/interface/main/web-goblin/service/web",
        "//app/interface/main/web-goblin/service/wechat",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
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
