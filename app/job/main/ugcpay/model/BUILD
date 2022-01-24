package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "model",
    srcs = [
        "db.go",
        "elec_db.go",
        "model.go",
    ],
    importpath = "go-common/app/job/main/ugcpay/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/log"],
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
    name = "model_test",
    srcs = ["model_test.go"],
    embed = [":model"],
    tags = ["automanaged"],
    deps = ["//vendor/github.com/smartystreets/goconvey/convey"],
)
