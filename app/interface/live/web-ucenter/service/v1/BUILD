package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "anchorTask.go",
        "capsule.go",
        "history.go",
        "user.go",
    ],
    importpath = "go-common/app/interface/live/web-ucenter/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/web-ucenter/api/http",
        "//app/interface/live/web-ucenter/api/http/v1:http",
        "//app/interface/live/web-ucenter/conf",
        "//app/interface/live/web-ucenter/dao",
        "//app/interface/live/web-ucenter/dao/capsule",
        "//app/interface/live/web-ucenter/dao/history",
        "//app/interface/live/web-ucenter/dao/user",
        "//app/interface/live/web-ucenter/model",
        "//app/service/live/room/api/liverpc/v2:liverpc",
        "//app/service/live/xrewardcenter/api/grpc/v1:grpc",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/sync/errgroup",
        "//vendor/github.com/pkg/errors",
        "//vendor/google.golang.org/grpc/status",
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
    name = "service_test",
    srcs = [
        "anchorTask_test.go",
        "history_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/live/web-ucenter/api/http/v1:http",
        "//app/interface/live/web-ucenter/conf",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
