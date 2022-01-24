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
    deps = ["//app/interface/main/dm2/model"],
)

go_library(
    name = "http",
    srcs = [
        "dm.go",
        "filter.go",
        "history.go",
        "http.go",
        "manage.go",
        "mask.go",
        "member.go",
        "subtitle.go",
        "thumbup.go",
        "view.go",
    ],
    importpath = "go-common/app/interface/main/dm2/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/dm2/conf",
        "//app/interface/main/dm2/model",
        "//app/interface/main/dm2/model/oplog",
        "//app/interface/main/dm2/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/middleware/antispam",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/ip",
        "//library/net/metadata",
        "//library/time",
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
