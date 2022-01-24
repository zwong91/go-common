package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "dM.go",
        "dmcheck.go",
        "getCheckInfo.go",
        "getDMConf.go",
        "limit.go",
        "send.go",
    ],
    importpath = "go-common/app/service/live/live-dm/service/v1",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/av/api/liverpc/v1:liverpc",
        "//app/service/live/banned_service/api/liverpc/v1:liverpc",
        "//app/service/live/live-dm/api/grpc/v1:grpc",
        "//app/service/live/live-dm/conf",
        "//app/service/live/live-dm/dao",
        "//app/service/live/live-dm/model",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
        "//vendor/github.com/ipipdotnet/ipdb-go",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/satori/go.uuid:go_uuid",
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
        "dm_test.go",
        "dmcheck_test.go",
        "getCheckInfo_test.go",
        "getDMConf_test.go",
        "limit_test.go",
        "send_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/live-dm/api/grpc/v1:grpc",
        "//app/service/live/live-dm/conf",
        "//app/service/live/live-dm/dao",
        "//library/ecode",
        "//library/net/metadata",
    ],
)
