package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "region",
    srcs = ["region.go"],
    importpath = "go-common/app/interface/main/app-show/model/region",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-show/model",
        "//app/interface/main/app-show/model/activity",
        "//app/interface/main/app-show/model/banner",
        "//app/interface/main/app-show/model/recommend",
        "//app/interface/main/app-show/model/tag",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/openplatform/pgc-season/api/grpc/season/v1:season",
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
