package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "model",
    srcs = ["const.go"],
    importpath = "go-common/app/interface/main/app-card/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-card/model/card/audio",
        "//app/interface/main/app-card/model/card/live",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/model/archive",
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
    srcs = [
        ":package-srcs",
        "//app/interface/main/app-card/model/bplus:all-srcs",
        "//app/interface/main/app-card/model/card:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
