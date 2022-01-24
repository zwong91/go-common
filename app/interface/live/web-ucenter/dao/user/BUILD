package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "user",
    srcs = ["dao.go"],
    importpath = "go-common/app/interface/live/web-ucenter/dao/user",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/web-ucenter/conf",
        "//app/interface/live/web-ucenter/dao",
        "//app/service/live/rankdb/api/liverpc/v1:liverpc",
        "//app/service/live/rc/api/liverpc/v1:liverpc",
        "//app/service/live/xuser/api/grpc/v1:grpc",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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

go_test(
    name = "user_test",
    srcs = ["dao_test.go"],
    embed = [":user"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/live/web-ucenter/conf",
        "//app/interface/live/web-ucenter/dao",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
