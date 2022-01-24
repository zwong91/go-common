package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "mail.go",
        "mng_auth.go",
        "mng_role.go",
        "mng_user.go",
        "package.go",
        "publish.go",
        "service.go",
    ],
    importpath = "go-common/app/admin/main/macross/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/macross/conf",
        "//app/admin/main/macross/dao",
        "//app/admin/main/macross/dao/oss",
        "//app/admin/main/macross/model/mail",
        "//app/admin/main/macross/model/manager",
        "//app/admin/main/macross/model/package",
        "//app/admin/main/macross/model/publish",
        "//app/admin/main/macross/tools",
        "//library/log",
        "//vendor/gopkg.in/gomail.v2:gomail_v2",
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
