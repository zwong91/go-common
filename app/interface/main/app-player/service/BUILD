package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["service_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-player/conf",
        "//app/interface/main/app-player/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "bnj.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/app-player/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-player/conf",
        "//app/interface/main/app-player/dao",
        "//app/interface/main/app-player/dao/account",
        "//app/interface/main/app-player/dao/archive",
        "//app/interface/main/app-player/dao/resource",
        "//app/interface/main/app-player/dao/ugcpay",
        "//app/interface/main/app-player/model",
        "//app/interface/main/app-player/model/archive",
        "//app/service/main/account/api",
        "//app/service/main/ugcpay/api/grpc/v1:grpc",
        "//library/ecode",
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
