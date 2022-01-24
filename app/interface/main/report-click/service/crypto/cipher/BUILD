package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "cipher_test",
    srcs = ["ecb_test.go"],
    embed = [":cipher"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/report-click/service/crypto/padding",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "cipher",
    srcs = ["ecb.go"],
    importpath = "go-common/app/interface/main/report-click/service/crypto/cipher",
    tags = ["automanaged"],
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
