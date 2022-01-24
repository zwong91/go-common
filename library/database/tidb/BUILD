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
    name = "tidb",
    srcs = [
        "discovery.go",
        "node_proc.go",
        "sql.go",
        "tidb.go",
    ],
    importpath = "go-common/library/database/tidb",
    tags = ["automanaged"],
    deps = [
        "//library/conf/env",
        "//library/log",
        "//library/naming",
        "//library/naming/discovery",
        "//library/net/netutil/breaker",
        "//library/net/trace",
        "//library/stat",
        "//library/time",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/pkg/errors",
    ],
)

go_test(
    name = "tidb_test",
    srcs = [
        "sql_test.go",
        "tidb_test.go",
    ],
    embed = [":tidb"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/net/netutil/breaker",
        "//library/time",
    ],
)
