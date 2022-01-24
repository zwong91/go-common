package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "view",
    srcs = [
        "access.go",
        "auth.go",
        "contain.go",
        "empty_arc.go",
        "service.go",
        "view.go",
    ],
    importpath = "go-common/app/interface/main/tv/service/view",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/dao/account",
        "//app/interface/main/tv/dao/archive",
        "//app/interface/main/tv/dao/cms",
        "//app/interface/main/tv/dao/favorite",
        "//app/interface/main/tv/dao/upper",
        "//app/interface/main/tv/model",
        "//app/interface/main/tv/model/upper",
        "//app/interface/main/tv/model/view",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//library/ecode",
        "//library/log",
        "//library/stat/prom",
        "//library/sync/errgroup",
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
