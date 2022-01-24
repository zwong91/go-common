package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "clean_cache.go",
        "service.go",
        "syn_auth.go",
        "sync_all.go",
    ],
    importpath = "go-common/app/job/main/passport-auth/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/passport-auth/conf",
        "//app/job/main/passport-auth/dao",
        "//app/job/main/passport-auth/model",
        "//app/service/main/passport-auth/rpc/client",
        "//library/log",
        "//library/queue/databus",
        "//library/queue/databus/databusutil",
        "//library/time",
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

go_test(
    name = "service_test",
    srcs = [
        "clean_cache_test.go",
        "service_test.go",
        "syn_auth_test.go",
        "sync_all_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/passport-auth/conf",
        "//app/job/main/passport-auth/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
