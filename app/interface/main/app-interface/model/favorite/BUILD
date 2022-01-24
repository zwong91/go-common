package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "favorite",
    srcs = [
        "folder.go",
        "result.go",
    ],
    importpath = "go-common/app/interface/main/app-interface/model/favorite",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/model",
        "//app/interface/main/app-interface/model/audio",
        "//app/interface/main/app-interface/model/bplus",
        "//app/interface/main/app-interface/model/sp",
        "//app/interface/main/app-interface/model/topic",
        "//app/interface/openplatform/article/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
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
