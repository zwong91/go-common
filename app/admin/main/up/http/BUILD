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
        "//app/admin/main/up/conf",
        "//vendor/github.com/davecgh/go-spew/spew",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "http",
    srcs = [
        "command.go",
        "http.go",
        "http_base.go",
        "signup.go",
        "up_data.go",
        "up_trend.go",
        "upcrm.go",
        "upload.go",
    ],
    importpath = "go-common/app/admin/main/up/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/conf",
        "//app/admin/main/up/dao/upcrm",
        "//app/admin/main/up/model",
        "//app/admin/main/up/model/datamodel",
        "//app/admin/main/up/model/signmodel",
        "//app/admin/main/up/model/upcrmmodel",
        "//app/admin/main/up/service",
        "//app/admin/main/up/util",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/http/blademaster/middleware/verify",
        "//vendor/github.com/siddontang/go-mysql/mysql",
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
