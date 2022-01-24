package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "creative",
    srcs = [
        "archive_material.go",
        "dao.cache.go",
        "dao.go",
        "mc.cache.go",
        "operation.go",
        "type.go",
        "version.go",
    ],
    importpath = "go-common/app/interface/main/creative/dao/creative",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/archive",
        "//app/interface/main/creative/model/operation",
        "//app/interface/main/creative/model/version",
        "//library/cache",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/log",
        "//library/stat/prom",
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
    name = "creative_test",
    srcs = [
        "dao_test.go",
        "material_test.go",
        "operation_test.go",
        "type_test.go",
        "version_test.go",
    ],
    embed = [":creative"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
