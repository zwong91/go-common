package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "activity.go",
        "column.go",
        "exchange.go",
        "http.go",
        "income.go",
        "join.go",
        "local.go",
        "newbie.go",
        "notice.go",
        "special_award.go",
        "web.go",
        "withdraw.go",
    ],
    importpath = "go-common/app/interface/main/growup/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/growup/conf",
        "//app/interface/main/growup/model",
        "//app/interface/main/growup/service",
        "//app/interface/main/growup/service/newbie",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/render",
        "//library/net/metadata",
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
