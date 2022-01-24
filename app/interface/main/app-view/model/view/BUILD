package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "view",
    srcs = [
        "bnj.go",
        "view.go",
    ],
    importpath = "go-common/app/interface/main/app-view/model/view",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-view/model",
        "//app/interface/main/app-view/model/act",
        "//app/interface/main/app-view/model/ad",
        "//app/interface/main/app-view/model/bangumi",
        "//app/interface/main/app-view/model/creative",
        "//app/interface/main/app-view/model/elec",
        "//app/interface/main/app-view/model/game",
        "//app/interface/main/app-view/model/live",
        "//app/interface/main/app-view/model/manager",
        "//app/interface/main/app-view/model/special",
        "//app/interface/main/app-view/model/tag",
        "//app/interface/main/dm2/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//app/service/main/resource/model",
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
