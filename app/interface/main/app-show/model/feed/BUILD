package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "feed",
    srcs = ["feed.go"],
    importpath = "go-common/app/interface/main/app-show/model/feed",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-card/model/card/live",
        "//app/interface/main/app-show/model",
        "//app/interface/main/app-show/model/card",
        "//app/interface/main/app-show/model/dislike",
        "//app/interface/main/app-show/model/tag",
        "//app/service/main/archive/api",
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
