package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "cursor.go",
        "http.go",
        "local.go",
        "reply.go",
        "reply_admin.go",
        "reply_record.go",
        "reply_report.go",
        "xreply.go",
    ],
    importpath = "go-common/app/interface/main/reply/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/reply/conf",
        "//app/interface/main/reply/model/adminlog",
        "//app/interface/main/reply/model/reply",
        "//app/interface/main/reply/model/xreply",
        "//app/interface/main/reply/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/metadata",
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
