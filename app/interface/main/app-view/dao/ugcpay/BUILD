package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "ugcpay_test",
    srcs = ["ugcpay_test.go"],
    embed = [":ugcpay"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-view/conf",
        "//app/interface/main/app-view/model/view",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "ugcpay",
    srcs = ["ugcpay.go"],
    importpath = "go-common/app/interface/main/app-view/dao/ugcpay",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-view/conf",
        "//app/interface/main/app-view/model/view",
        "//app/service/main/ugcpay/api/grpc/v1:grpc",
        "//vendor/github.com/pkg/errors",
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
