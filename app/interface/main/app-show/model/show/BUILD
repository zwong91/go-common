package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "show",
    srcs = ["show.go"],
    importpath = "go-common/app/interface/main/app-show/model/show",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-show/model",
        "//app/interface/main/app-show/model/activity",
        "//app/interface/main/app-show/model/bangumi",
        "//app/interface/main/app-show/model/banner",
        "//app/interface/main/app-show/model/live",
        "//app/interface/main/app-show/model/recommend",
        "//app/service/main/archive/api",
        "//app/service/openplatform/pgc-season/api/grpc/season/v1:season",
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
