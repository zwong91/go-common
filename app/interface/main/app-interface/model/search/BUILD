package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "search",
    srcs = [
        "from.go",
        "result.go",
    ],
    importpath = "go-common/app/interface/main/app-interface/model/search",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/model",
        "//app/interface/main/app-interface/model/bangumi",
        "//app/interface/main/app-interface/model/bplus",
        "//app/interface/main/app-interface/model/live",
        "//app/interface/main/app-interface/model/tag",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//app/service/openplatform/pgc-season/api/grpc/season/v1:season",
        "//library/log",
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
