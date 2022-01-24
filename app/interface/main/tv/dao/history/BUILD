package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "history_test",
    srcs = [
        "cursor_test.go",
        "dao_test.go",
    ],
    embed = [":history"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/history/model",
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model/history",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "history",
    srcs = [
        "cursor.go",
        "dao.go",
    ],
    importpath = "go-common/app/interface/main/tv/dao/history",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/history/model",
        "//app/interface/main/history/rpc/client",
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model/history",
        "//library/cache/memcache",
        "//library/log",
        "//library/net/metadata",
        "//library/stat/prom",
        "//vendor/github.com/pkg/errors",
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
