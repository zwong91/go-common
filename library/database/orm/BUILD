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
)

go_library(
    name = "orm",
    srcs = [
        "orm.go",
        "timestamp.go",
    ],
    importpath = "go-common/library/database/orm",
    tags = ["automanaged"],
    deps = [
        "//library/ecode",
        "//library/log",
        "//library/time",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/jinzhu/gorm",
    ],
)
