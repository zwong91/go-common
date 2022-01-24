load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

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
)

go_library(
    name = "sql",
    srcs = [
        "mysql.go",
        "sql.go",
    ],
    importpath = "go-common/library/database/sql",
    tags = ["automanaged"],
    deps = [
        "//library/ecode",
        "//library/log",
        "//library/net/netutil/breaker",
        "//library/net/trace",
        "//library/stat",
        "//library/time",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/pkg/errors",
    ],
)

go_test(
    name = "sql_test",
    srcs = ["mysql_test.go"],
    embed = [":sql"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/net/netutil/breaker",
        "//library/time",
    ],
)
