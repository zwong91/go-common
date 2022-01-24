package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dispatch_test",
    srcs = ["matcher_test.go"],
    embed = [":dispatch"],
    tags = ["automanaged"],
)

go_library(
    name = "dispatch",
    srcs = [
        "matcher.go",
        "min_heap.go",
        "sinaip.go",
    ],
    importpath = "go-common/app/service/live/broadcast-proxy/dispatch",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/broadcast-proxy/conf",
        "//app/service/live/broadcast-proxy/expr",
        "//library/log",
        "//vendor/github.com/ipipdotnet/ipdb-go",
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
