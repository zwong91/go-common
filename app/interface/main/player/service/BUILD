package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "archive_test.go",
        "player_test.go",
        "policy_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/player/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "archive.go",
        "player.go",
        "playurl.go",
        "policy.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/player/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/dm2/model",
        "//app/interface/main/dm2/rpc/client",
        "//app/interface/main/history/model",
        "//app/interface/main/history/rpc/client",
        "//app/interface/main/player/conf",
        "//app/interface/main/player/dao",
        "//app/interface/main/player/model",
        "//app/interface/main/tag/model",
        "//app/interface/main/tag/rpc/client",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/assist/model/assist",
        "//app/service/main/assist/rpc/client",
        "//app/service/main/location/model",
        "//app/service/main/location/rpc/client",
        "//app/service/main/resource/model",
        "//app/service/main/resource/rpc/client",
        "//app/service/main/ugcpay/api/grpc/v1:grpc",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
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
