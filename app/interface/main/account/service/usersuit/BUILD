package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "usersuit",
    srcs = [
        "invite.go",
        "pendant.go",
        "pendant_new.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/account/service/usersuit",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/interface/main/account/dao/usersuit",
        "//app/interface/main/account/dao/vip",
        "//app/interface/main/account/model",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//app/service/main/coin/api/gorpc",
        "//app/service/main/coin/model",
        "//app/service/main/member/api/gorpc",
        "//app/service/main/member/model",
        "//app/service/main/usersuit/model",
        "//app/service/main/usersuit/rpc/client",
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

go_test(
    name = "usersuit_test",
    srcs = [
        "invite_test.go",
        "pendant_test.go",
        "service_test.go",
    ],
    embed = [":usersuit"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/service/main/usersuit/model",
        "//app/service/main/usersuit/rpc/client",
        "//library/ecode",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
