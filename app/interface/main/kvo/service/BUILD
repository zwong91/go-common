package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["document_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/kvo/conf",
        "//app/interface/main/kvo/model/module",
        "//library/cache/memcache",
        "//library/container/pool",
        "//library/database/sql",
        "//library/net/netutil/breaker",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "document.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/kvo/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/kvo/conf",
        "//app/interface/main/kvo/dao",
        "//app/interface/main/kvo/model",
        "//app/interface/main/kvo/model/module",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
