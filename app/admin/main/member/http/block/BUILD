package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "block_test",
    srcs = ["http_test.go"],
    embed = [":block"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/block/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "block",
    srcs = [
        "block.go",
        "http.go",
    ],
    importpath = "go-common/app/admin/main/member/http/block",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/member/model/block",
        "//app/admin/main/member/service/block",
        "//library/ecode",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
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
