package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "data",
    srcs = [
        "service.go",
        "uparchiveinfo.go",
        "viewer.go",
    ],
    importpath = "go-common/app/admin/main/up/service/data",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/conf",
        "//app/admin/main/up/dao/data",
        "//app/admin/main/up/dao/tag",
        "//app/admin/main/up/model/datamodel",
        "//app/admin/main/up/service/cache",
        "//app/admin/main/up/util",
        "//app/interface/main/creative/model/tag",
        "//library/log",
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

go_test(
    name = "data_test",
    srcs = [
        "service_test.go",
        "uparchiveinfo_test.go",
        "viewer_test.go",
    ],
    embed = [":data"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/up/conf",
        "//app/admin/main/up/model/datamodel",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
