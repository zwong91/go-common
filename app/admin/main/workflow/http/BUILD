package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "activity.go",
        "business.go",
        "callback.go",
        "challenge.go",
        "event.go",
        "group.go",
        "http.go",
        "permission.go",
        "reply.go",
        "util.go",
    ],
    importpath = "go-common/app/admin/main/workflow/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/workflow/model",
        "//app/admin/main/workflow/model/param",
        "//app/admin/main/workflow/model/search",
        "//app/admin/main/workflow/service",
        "//library/conf/paladin",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/http/blademaster/render",
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
