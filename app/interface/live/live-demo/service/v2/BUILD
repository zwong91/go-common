package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["service_test.go"],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/live/live-demo/api/http/v2:http",
        "//app/interface/live/live-demo/conf",
        "//app/interface/live/live-demo/dao",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "foo.go",
        "foo2.go",
    ],
    importpath = "go-common/app/interface/live/live-demo/service/v2",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/live-demo/api/http/v2:http",
        "//app/interface/live/live-demo/conf",
        "//app/interface/live/live-demo/dao",
        "//app/service/live/room/api/liverpc/v2:liverpc",
        "//library/ecode",
        "//library/ecode/pb",
        "//library/log",
        "//library/net/metadata",
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
