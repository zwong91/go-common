package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "relations_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/relation/conf",
        "//app/admin/main/relation/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "relation_log.go",
        "relations.go",
        "service.go",
    ],
    importpath = "go-common/app/admin/main/relation/service",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/relation/conf",
        "//app/admin/main/relation/dao",
        "//app/admin/main/relation/model",
        "//app/service/main/relation/api",
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
