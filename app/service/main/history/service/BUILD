package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "history.go",
        "service.go",
        "user.go",
    ],
    importpath = "go-common/app/service/main/history/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/history/api/grpc",
        "//app/service/main/history/conf",
        "//app/service/main/history/dao",
        "//app/service/main/history/model",
        "//library/ecode",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//library/sync/pipeline",
        "//library/sync/pipeline/fanout",
        "//vendor/golang.org/x/sync/singleflight",
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
    name = "service_test",
    srcs = [
        "history_test.go",
        "service_test.go",
        "user_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/history/api/grpc",
        "//app/service/main/history/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
