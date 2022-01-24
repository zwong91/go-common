package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "http_test",
    srcs = [
        "adv_test.go",
        "http_test.go",
    ],
    embed = [":http"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "http",
    srcs = [
        "adv.go",
        "assist.go",
        "dm.go",
        "filter.go",
        "http.go",
        "report.go",
        "up.go",
    ],
    importpath = "go-common/app/interface/main/dm/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/dm/conf",
        "//app/interface/main/dm/model",
        "//app/interface/main/dm/service",
        "//app/interface/main/dm2/model",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/antispam",
        "//library/net/http/blademaster/middleware/auth",
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
