package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "authority.go",
        "dao.go",
        "parameter.go",
        "score.go",
        "statistics.go",
    ],
    importpath = "go-common/app/admin/main/up-rating/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up-rating/conf",
        "//app/admin/main/up-rating/model",
        "//library/database/sql",
        "//library/log",
        "//library/xstr",
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
    srcs = [
        ":package-srcs",
        "//app/admin/main/up-rating/dao/global:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "dao_test",
    srcs = [
        "authority_test.go",
        "dao_test.go",
        "parameter_test.go",
        "score_test.go",
        "statistics_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/up-rating/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
