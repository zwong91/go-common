package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["recommend_test.go"],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/recommend/api/grpc/v1:grpc",
        "//app/service/live/recommend/internal/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = ["recommend.go"],
    importpath = "go-common/app/service/live/recommend/internal/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/recommend/api/grpc/v1:grpc",
        "//app/service/live/recommend/internal/conf",
        "//app/service/live/recommend/internal/dao",
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
