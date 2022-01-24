package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "infoc.go",
        "package.go",
        "service.go",
    ],
    importpath = "go-common/app/job/live/gift/internal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/live/gift/api",
        "//app/job/live/gift/internal/conf",
        "//app/job/live/gift/internal/dao",
        "//app/job/live/gift/internal/model",
        "//library/log",
        "//library/log/infoc",
        "//library/queue/databus",
        "//library/sync/errgroup",
        "//vendor/github.com/golang/protobuf/ptypes/empty",
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
        "infoc_test.go",
        "package_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/job/live/gift/internal/conf",
        "//app/job/live/gift/internal/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
