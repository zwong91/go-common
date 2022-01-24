package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "auth.go",
        "dao.go",
        "mc.go",
        "secret.go",
        "user.go",
    ],
    importpath = "go-common/app/interface/main/passport-login/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/passport-login/conf",
        "//app/interface/main/passport-login/model",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/log",
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
        "auth_test.go",
        "dao_test.go",
        "mc_test.go",
        "secret_test.go",
        "user_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/passport-login/conf",
        "//app/interface/main/passport-login/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
