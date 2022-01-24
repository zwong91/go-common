package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "newcomer",
    srcs = [
        "api.go",
        "dao.go",
        "db.go",
    ],
    importpath = "go-common/app/job/main/creative/dao/newcomer",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/api",
        "//app/job/main/creative/conf",
        "//app/job/main/creative/model",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
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

go_test(
    name = "newcomer_test",
    srcs = [
        "api_test.go",
        "dao_test.go",
        "db_test.go",
    ],
    embed = [":newcomer"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/creative/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
