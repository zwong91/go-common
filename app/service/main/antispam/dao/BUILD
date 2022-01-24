package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "cond_test.go",
        "dao_test.go",
        "keyword_test.go",
        "mysql_test.go",
        "redis_test.go",
        "regexp_test.go",
        "rule_test.go",
        "transaction_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/antispam/conf",
        "//app/service/main/antispam/model",
        "//app/service/main/antispam/util",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "cond.go",
        "dao.go",
        "keyword.go",
        "mysql.go",
        "redis.go",
        "regexp.go",
        "rule.go",
        "transaction.go",
    ],
    importpath = "go-common/app/service/main/antispam/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/antispam/conf",
        "//app/service/main/antispam/model",
        "//app/service/main/antispam/util",
        "//library/cache/redis",
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
