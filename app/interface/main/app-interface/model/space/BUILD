package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "space",
    srcs = [
        "contribute.go",
        "mine.go",
        "space.go",
    ],
    importpath = "go-common/app/interface/main/app-interface/model/space",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/model",
        "//app/interface/main/app-interface/model/audio",
        "//app/interface/main/app-interface/model/bangumi",
        "//app/interface/main/app-interface/model/bplus",
        "//app/interface/main/app-interface/model/community",
        "//app/interface/main/app-interface/model/elec",
        "//app/interface/main/app-interface/model/favorite",
        "//app/interface/main/tag/model",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/api",
        "//app/service/main/account/model",
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
