package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "seq_test",
    srcs = ["seq_test.go"],
    embed = [":seq"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-wall/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "seq",
    srcs = ["seq.go"],
    importpath = "go-common/app/interface/main/app-wall/dao/seq",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-wall/conf",
        "//app/service/main/seq-server/model",
        "//app/service/main/seq-server/rpc/client",
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
