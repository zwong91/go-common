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
        "//app/service/main/archive/conf",
        "//app/service/main/archive/service",
        "//library/net/http/blademaster",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "http",
    srcs = [
        "archive.go",
        "http.go",
        "local.go",
        "region.go",
        "shot.go",
        "stat.go",
        "upper.go",
        "video.go",
    ],
    importpath = "go-common/app/service/main/archive/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/archive/conf",
        "//app/service/main/archive/model/archive",
        "//app/service/main/archive/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/verify",
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
