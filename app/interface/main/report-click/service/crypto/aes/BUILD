package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "aes_test",
    srcs = ["aes_test.go"],
    embed = [":aes"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/report-click/service/crypto/padding",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "aes",
    srcs = ["aes.go"],
    importpath = "go-common/app/interface/main/report-click/service/crypto/aes",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/report-click/service/crypto/cipher",
        "//app/interface/main/report-click/service/crypto/padding",
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
