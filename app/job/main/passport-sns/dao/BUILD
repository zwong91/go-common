package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "aso_test.go",
        "dao_test.go",
        "mc_test.go",
        "qq_test.go",
        "sns_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/passport-sns/conf",
        "//app/job/main/passport-sns/model",
        "//library/database/sql",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "aso.go",
        "dao.go",
        "mc.go",
        "qq.go",
        "sns.go",
    ],
    importpath = "go-common/app/job/main/passport-sns/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/passport-sns/conf",
        "//app/job/main/passport-sns/model",
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
