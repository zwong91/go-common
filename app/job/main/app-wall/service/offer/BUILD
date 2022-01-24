package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "offer_test",
    srcs = [
        "active_test.go",
        "retry_test.go",
        "service_test.go",
    ],
    embed = [":offer"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/app-wall/conf",
        "//app/job/main/app-wall/model/offer",
        "//vendor/github.com/bsm/sarama-cluster",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "offer",
    srcs = [
        "active.go",
        "retry.go",
        "service.go",
    ],
    importpath = "go-common/app/job/main/app-wall/service/offer",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/app-wall/conf",
        "//app/job/main/app-wall/dao/offer",
        "//app/job/main/app-wall/model",
        "//app/job/main/app-wall/model/offer",
        "//library/log",
        "//vendor/github.com/Shopify/sarama",
        "//vendor/github.com/bsm/sarama-cluster",
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
