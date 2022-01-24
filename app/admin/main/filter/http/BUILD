package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "http_test",
    srcs = ["http_test.go"],
    embed = [":http"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/filter/conf",
        "//app/admin/main/filter/model",
        "//app/admin/main/filter/service",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "http",
    srcs = [
        "ai.go",
        "area.go",
        "filter.go",
        "http.go",
        "key.go",
        "white.go",
    ],
    importpath = "go-common/app/admin/main/filter/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/filter/conf",
        "//app/admin/main/filter/model",
        "//app/admin/main/filter/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/http/blademaster/middleware/verify",
        "//library/xstr",
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
