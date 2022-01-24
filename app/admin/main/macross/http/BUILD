package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "http.go",
        "local.go",
        "mail.go",
        "manager.go",
        "mng_auth.go",
        "mng_role.go",
        "mng_user.go",
        "package.go",
        "publish.go",
    ],
    importpath = "go-common/app/admin/main/macross/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/macross/conf",
        "//app/admin/main/macross/model/mail",
        "//app/admin/main/macross/model/package",
        "//app/admin/main/macross/model/publish",
        "//app/admin/main/macross/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/http/blademaster/middleware/verify",
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
