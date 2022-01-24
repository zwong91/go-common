package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "player",
    srcs = ["service.go"],
    importpath = "go-common/app/interface/main/app-intl/service/player",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-intl/conf",
        "//app/interface/main/app-intl/dao/account",
        "//app/interface/main/app-intl/dao/archive",
        "//app/interface/main/app-intl/dao/player",
        "//app/interface/main/app-intl/dao/resource",
        "//app/interface/main/app-intl/dao/ugcpay",
        "//app/interface/main/app-intl/model",
        "//app/interface/main/app-intl/model/player",
        "//app/interface/main/app-intl/model/player/archive",
        "//app/service/main/account/model",
        "//app/service/main/ugcpay/api/grpc/v1:grpc",
        "//library/ecode",
        "//library/log",
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
