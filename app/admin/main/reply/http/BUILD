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
        "//app/admin/main/reply/conf",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
    ],
)

go_library(
    name = "http",
    srcs = [
        "action.go",
        "business.go",
        "emoji.go",
        "emoji_package.go",
        "http.go",
        "log.go",
        "monitor.go",
        "notice.go",
        "reply.go",
        "report.go",
        "subject.go",
    ],
    importpath = "go-common/app/admin/main/reply/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/reply/conf",
        "//app/admin/main/reply/model",
        "//app/admin/main/reply/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/ip",
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
