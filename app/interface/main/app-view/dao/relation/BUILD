package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "relation_test",
    srcs = ["dao_test.go"],
    embed = [":relation"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-view/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "relation",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/main/app-view/dao/relation",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-view/conf",
        "//app/service/main/relation/model",
        "//app/service/main/relation/rpc/client",
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
