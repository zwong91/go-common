package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "history_test.go",
        "service_test.go",
        "toview_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/history/conf",
        "//app/interface/main/history/model",
        "//library/cache/redis",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "history.go",
        "history_service.go",
        "rpc.go",
        "service.go",
        "toview.go",
    ],
    importpath = "go-common/app/interface/main/history/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/history/conf",
        "//app/interface/main/history/dao/history",
        "//app/interface/main/history/dao/toview",
        "//app/interface/main/history/model",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/favorite/api/gorpc",
        "//app/service/main/favorite/model",
        "//app/service/main/history/api/grpc",
        "//app/service/main/history/model",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus/report",
        "//library/sync/pipeline/fanout",
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
