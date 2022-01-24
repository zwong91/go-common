package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "livewallet_test.go",
        "paycenter_test.go",
        "userext_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/live/app-room/conf",
        "//app/interface/live/app-room/model",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "api.go",
        "dao.go",
        "livewallet.go",
        "paycenter.go",
        "userext.go",
    ],
    importpath = "go-common/app/interface/live/app-room/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-room/conf",
        "//app/interface/live/app-room/model",
        "//app/service/live/userext/api/liverpc",
        "//app/service/live/userext/api/liverpc/v0:liverpc",
        "//app/service/live/xuserex/api/grpc/v1:grpc",
        "//library/cache",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/rpc/liverpc",
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
        "//app/interface/live/app-room/dao/av:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
