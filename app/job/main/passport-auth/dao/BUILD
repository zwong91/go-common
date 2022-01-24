package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "clean_cache.go",
        "cookie.go",
        "dao.go",
        "refresh.go",
        "token.go",
    ],
    importpath = "go-common/app/job/main/passport-auth/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/passport-auth/conf",
        "//app/job/main/passport-auth/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
    name = "dao_test",
    srcs = [
        "clean_cache_test.go",
        "cookie_test.go",
        "dao_test.go",
        "refresh_test.go",
        "token_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/passport-auth/conf",
        "//app/job/main/passport-auth/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
