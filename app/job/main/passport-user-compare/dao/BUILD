package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "duplicate.go",
        "origin.go",
        "secret.go",
        "user.go",
        "wechat.go",
    ],
    importpath = "go-common/app/job/main/passport-user-compare/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/passport-user-compare/conf",
        "//app/job/main/passport-user-compare/model",
        "//library/database/sql",
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
        "dao_test.go",
        "duplicate_test.go",
        "origin_test.go",
        "secret_test.go",
        "user_test.go",
        "wechat_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/passport-user-compare/conf",
        "//app/job/main/passport-user-compare/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
