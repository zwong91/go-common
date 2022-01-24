package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "model",
    srcs = [
        "abserver.go",
        "account.go",
        "archive.go",
        "article.go",
        "baidu.go",
        "bigdata.go",
        "bnj.go",
        "coin.go",
        "elec.go",
        "feedback.go",
        "help.go",
        "icon.go",
        "nav.go",
        "online.go",
        "reply.go",
        "search.go",
        "shop.go",
        "tag.go",
        "web.go",
        "wechat.go",
    ],
    importpath = "go-common/app/interface/main/web/model",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/dm2/model",
        "//app/interface/main/reply/model/reply",
        "//app/interface/main/tag/model",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/broadcast/api/grpc/v1:grpc",
        "//app/service/main/ugcpay/api/grpc/v1:grpc",
        "//library/time",
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
