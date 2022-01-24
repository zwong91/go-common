package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "base.go",
        "dao.go",
        "paramter.go",
        "past.go",
        "rating.go",
        "statistics.go",
        "task_status.go",
    ],
    importpath = "go-common/app/job/main/up-rating/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/up-rating/conf",
        "//app/job/main/up-rating/model",
        "//library/database/sql",
        "//library/log",
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
    name = "dao_test",
    srcs = [
        "base_test.go",
        "dao_test.go",
        "paramter_test.go",
        "past_test.go",
        "rating_test.go",
        "statistics_test.go",
        "task_status_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/up-rating/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
