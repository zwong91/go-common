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
        "//app/interface/main/favorite/conf",
        "//app/interface/main/favorite/service",
        "//library/net/http/blademaster",
    ],
)

go_library(
    name = "http",
    srcs = [
        "http.go",
        "topic.go",
        "video.go",
    ],
    importpath = "go-common/app/interface/main/favorite/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/favorite/conf",
        "//app/interface/main/favorite/model",
        "//app/interface/main/favorite/service",
        "//library/ecode",
        "//library/log",
        "//library/log/anticheat",
        "//library/log/infoc",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/antispam",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/supervisor",
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
