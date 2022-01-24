package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "model_test",
    srcs = ["regular_test.go"],
    embed = [":model"],
    tags = ["automanaged"],
    deps = ["//vendor/github.com/smartystreets/goconvey/convey"],
)

go_library(
    name = "model",
    srcs = [
        "app_info.go",
        "card.go",
        "coupon.go",
        "enum.go",
        "geetest.go",
        "http.go",
        "invite.go",
        "model.go",
        "notice.go",
        "pendant.go",
        "realname.go",
        "regular.go",
        "relation.go",
        "reply.go",
        "usersuit.go",
        "vip.go",
        "vip_outer.go",
        "vip_welfare.go",
        "web_info.go",
    ],
    importpath = "go-common/app/interface/main/account/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//app/service/main/coupon/model",
        "//app/service/main/relation/model",
        "//app/service/main/usersuit/model",
        "//app/service/main/vip/api",
        "//app/service/main/vip/model",
        "//library/net/http/blademaster",
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
