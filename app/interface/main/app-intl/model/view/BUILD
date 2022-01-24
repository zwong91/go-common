package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "view",
    srcs = ["view.go"],
    importpath = "go-common/app/interface/main/app-intl/model/view",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-intl/model",
        "//app/interface/main/app-intl/model/bangumi",
        "//app/interface/main/app-intl/model/manager",
        "//app/interface/main/app-intl/model/tag",
        "//app/interface/main/dm2/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
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
