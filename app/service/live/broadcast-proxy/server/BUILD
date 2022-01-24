package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "server",
    srcs = [
        "dispatch.go",
        "http.go",
        "proxy.go",
        "zk.go",
    ],
    importpath = "go-common/app/service/live/broadcast-proxy/server",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/broadcast-proxy/conf",
        "//app/service/live/broadcast-proxy/dispatch",
        "//app/service/live/broadcast-proxy/grocery",
        "//library/log",
        "//vendor/github.com/ipipdotnet/ipdb-go",
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
