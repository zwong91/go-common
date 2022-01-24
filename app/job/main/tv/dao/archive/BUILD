package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "archive",
    srcs = [
        "archive.go",
        "dao.go",
    ],
    importpath = "go-common/app/job/main/tv/dao/archive",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/service/main/archive/api",
        "//library/cache/memcache",
        "//vendor/github.com/pkg/errors",
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
    name = "archive_test",
    srcs = [
        "archive_test.go",
        "dao_test.go",
    ],
    embed = [":archive"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/service/main/archive/api",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
