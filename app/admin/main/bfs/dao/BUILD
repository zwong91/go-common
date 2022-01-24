package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "http.go",
        "zk.go",
    ],
    importpath = "go-common/app/admin/main/bfs/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/bfs/conf",
        "//app/admin/main/bfs/model",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
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

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "http_test.go",
        "zk_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/bfs/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
