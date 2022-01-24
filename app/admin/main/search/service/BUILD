load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

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
    name = "service_test",
    srcs = [
        "archive_test.go",
        "log_test.go",
        "mng_v2_test.go",
        "service_test.go",
        "update_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/search/conf",
        "//app/admin/main/search/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "archive.go",
        "log.go",
        "mng.go",
        "mng_v2.go",
        "query.go",
        "service.go",
        "update.go",
        "upsert.go",
    ],
    importpath = "go-common/app/admin/main/search/service",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/search/conf",
        "//app/admin/main/search/dao",
        "//app/admin/main/search/model",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup",
    ],
)
