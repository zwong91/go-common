package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "content_ext_test.go",
        "service_test.go",
        "upload_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/up/conf",
        "//library/net/http/blademaster",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "content_ext.go",
        "service.go",
        "upload.go",
    ],
    importpath = "go-common/app/admin/main/up/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/conf",
        "//app/admin/main/up/dao",
        "//app/admin/main/up/dao/global",
        "//app/admin/main/up/dao/manager",
        "//app/admin/main/up/service/upcrmservice",
        "//app/admin/main/up/util",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/render",
        "//library/stat/prom",
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
    srcs = [
        ":package-srcs",
        "//app/admin/main/up/service/cache:all-srcs",
        "//app/admin/main/up/service/data:all-srcs",
        "//app/admin/main/up/service/upcrmservice:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
