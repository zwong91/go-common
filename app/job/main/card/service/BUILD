package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "service.go",
        "vip.go",
    ],
    importpath = "go-common/app/job/main/card/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/card/conf",
        "//app/job/main/card/dao",
        "//app/job/main/card/model",
        "//app/service/main/card/api/grpc/v1:grpc",
        "//app/service/main/card/model",
        "//app/service/main/vip/model",
        "//library/log",
        "//library/queue/databus",
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
    srcs = ["vip_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/card/conf",
        "//app/job/main/card/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
