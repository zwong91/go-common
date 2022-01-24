package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "doc.go",
        "index.go",
        "relation.go",
    ],
    importpath = "go-common/app/interface/live/app-interface/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-interface/api/http/v1:http",
        "//app/interface/live/app-interface/conf",
        "//app/interface/live/app-interface/dao",
        "//app/interface/live/app-interface/service/v1/relation",
        "//app/service/live/av/api/liverpc/v1:liverpc",
        "//app/service/live/live_user/api/liverpc/v1:liverpc",
        "//app/service/live/relation/api/liverpc/v1:liverpc",
        "//app/service/live/relation/api/liverpc/v2:liverpc",
        "//app/service/live/room/api/liverpc/v1:liverpc",
        "//app/service/live/room/api/liverpc/v2:liverpc",
        "//app/service/live/room_ex/api/liverpc/v1:liverpc",
        "//app/service/live/third_api/bvc",
        "//app/service/live/userext/api/liverpc/v1:liverpc",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/net/rpc/liverpc",
        "//library/net/rpc/liverpc/context",
        "//library/sync/errgroup",
        "//library/xstr",
        "//vendor/github.com/bitly/go-simplejson",
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
    srcs = [
        ":package-srcs",
        "//app/interface/live/app-interface/service/v1/app_conf:all-srcs",
        "//app/interface/live/app-interface/service/v1/relation:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "service_test",
    srcs = ["index_test.go"],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/live/app-interface/api/http/v1:http",
        "//app/interface/live/app-interface/conf",
        "//library/net/http/blademaster",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
