package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "mock",
    srcs = ["monkey_dao.go"],
    importpath = "go-common/app/admin/main/apm/dao/mock",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/apm/dao",
        "//app/admin/main/apm/model/canal",
        "//app/admin/main/apm/model/ecode",
        "//app/admin/main/apm/model/need",
        "//app/admin/main/apm/model/pprof",
        "//app/admin/main/apm/model/ut",
        "//vendor/github.com/bouk/monkey",
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
