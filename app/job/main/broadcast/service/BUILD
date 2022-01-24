package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "comet.go",
        "push.go",
        "room.go",
        "service.go",
    ],
    importpath = "go-common/app/job/main/broadcast/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/broadcast/api/grpc/v1:grpc",
        "//app/job/main/broadcast/conf",
        "//app/service/main/broadcast/libs/bytes",
        "//app/service/main/broadcast/model",
        "//library/conf/env",
        "//library/log",
        "//library/naming",
        "//library/naming/discovery",
        "//library/queue/databus",
        "//library/time",
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
    srcs = ["service_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/broadcast/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
