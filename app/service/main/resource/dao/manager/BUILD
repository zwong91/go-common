package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "manager_test",
    srcs = [
        "dao_test.go",
        "special_test.go",
    ],
    embed = [":manager"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/resource/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "manager",
    srcs = [
        "dao.go",
        "special.go",
    ],
    importpath = "go-common/app/service/main/resource/dao/manager",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/resource/api/v1:api",
        "//app/service/main/resource/conf",
        "//library/database/sql",
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
