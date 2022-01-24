package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "batch_query_test.go",
        "cloud2local_test.go",
        "parse_diff_log_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/passport-game-data/conf",
        "//app/job/main/passport-game-data/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "batch_query.go",
        "cloud2local.go",
        "compare_only.go",
        "diff.go",
        "init_cloud.go",
        "local2cloud.go",
        "log_and_fix.go",
        "parse_diff_log.go",
        "service.go",
        "stat.go",
    ],
    importpath = "go-common/app/job/main/passport-game-data/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/passport-game-data/conf",
        "//app/job/main/passport-game-data/dao",
        "//app/job/main/passport-game-data/model",
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
