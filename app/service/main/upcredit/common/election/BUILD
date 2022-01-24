package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "election_test",
    srcs = ["example_test.go"],
    embed = [":election"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//library/log"],
)

go_library(
    name = "election",
    srcs = ["election.go"],
    importpath = "go-common/app/service/main/upcredit/common/election",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//vendor/github.com/samuel/go-zookeeper/zk",
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
