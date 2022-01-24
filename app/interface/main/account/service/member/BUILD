package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "member_test",
    srcs = [
        "app_info_test.go",
        "notice_test.go",
        "remotelogin_test.go",
        "reply_test.go",
        "service_test.go",
        "web_info_test.go",
    ],
    embed = [":member"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/interface/main/account/model",
        "//app/service/main/secure/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "member",
    srcs = [
        "app_info.go",
        "medal.go",
        "member.go",
        "notice.go",
        "remotelogin.go",
        "reply.go",
        "service.go",
        "sudo.go",
        "web_info.go",
    ],
    importpath = "go-common/app/interface/main/account/service/member",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/account/conf",
        "//app/interface/main/account/dao/account",
        "//app/interface/main/account/dao/passport",
        "//app/interface/main/account/dao/reply",
        "//app/interface/main/account/model",
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/rpc/client",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/coin/api/gorpc",
        "//app/service/main/coin/model",
        "//app/service/main/filter/model/rpc",
        "//app/service/main/filter/rpc/client",
        "//app/service/main/location/model",
        "//app/service/main/location/rpc/client",
        "//app/service/main/member/api/gorpc",
        "//app/service/main/member/model",
        "//app/service/main/passport/model",
        "//app/service/main/passport/rpc/client",
        "//app/service/main/secure/model",
        "//app/service/main/secure/rpc/client",
        "//app/service/main/up/api/gorpc",
        "//app/service/main/usersuit/model",
        "//app/service/main/usersuit/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus",
        "//library/queue/databus/report",
        "//library/time",
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
