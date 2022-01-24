load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
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

go_library(
    name = "http",
    srcs = [
        "archive.go",
        "archive_update.go",
        "copyright.go",
        "http.go",
        "log.go",
        "log_delete.go",
        "mng.go",
        "mng_v2.go",
        "query.go",
        "upsert.go",
    ],
    importpath = "go-common/app/admin/main/search/http",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/search/conf",
        "//app/admin/main/search/dao",
        "//app/admin/main/search/model",
        "//app/admin/main/search/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
    ],
)
