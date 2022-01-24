package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "guard_test.go",
        "node_test.go",
        "nodes_test.go",
        "registry_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/infra/discovery/conf",
        "//app/infra/discovery/model",
        "//library/ecode",
        "//library/net/http/blademaster",
        "//library/net/netutil/breaker",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "guard.go",
        "node.go",
        "nodes.go",
        "registry.go",
    ],
    importpath = "go-common/app/infra/discovery/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/infra/discovery/conf",
        "//app/infra/discovery/model",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/sync/errgroup",
        "//library/xstr",
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
