package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "rpc_test.go",
        "service_test.go",
        "topic_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/favorite/conf",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "rpc.go",
        "service.go",
        "topic.go",
        "video.go",
    ],
    importpath = "go-common/app/interface/main/favorite/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/favorite/conf",
        "//app/interface/main/favorite/dao/music",
        "//app/interface/main/favorite/dao/topic",
        "//app/interface/main/favorite/dao/video",
        "//app/interface/main/favorite/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/favorite/api",
        "//app/service/main/favorite/model",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
        "//library/time",
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
