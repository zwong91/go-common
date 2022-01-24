package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "canal.go",
        "hbase.go",
        "instance.go",
        "master.go",
        "target.go",
        "tidb_check.go",
        "tidb_data.go",
        "tidb_instance.go",
        "tidb_proc.go",
    ],
    importpath = "go-common/app/infra/canal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/infra/canal/conf",
        "//app/infra/canal/dao",
        "//app/infra/canal/infoc",
        "//app/infra/canal/model",
        "//app/infra/canal/service/reader",
        "//library/conf/env",
        "//library/database/hbase.v2:hbase_v2",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/netutil",
        "//library/queue/databus",
        "//library/stat/prom",
        "//vendor/github.com/juju/errors",
        "//vendor/github.com/pingcap/tidb-tools/tidb_binlog/slave_binlog_proto/go-binlog",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/siddontang/go-mysql/canal",
        "//vendor/github.com/siddontang/go-mysql/client",
        "//vendor/github.com/siddontang/go-mysql/mysql",
        "//vendor/github.com/siddontang/go-mysql/replication",
        "//vendor/github.com/tsuna/gohbase/hrpc",
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
    srcs = [
        ":package-srcs",
        "//app/infra/canal/service/reader:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "service_test",
    srcs = ["tidb_data_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/infra/canal/model",
        "//vendor/github.com/pingcap/tidb-tools/tidb_binlog/slave_binlog_proto/go-binlog",
    ],
)
