package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "invite_test.go",
        "medal_test.go",
        "notify_test.go",
        "pendant_test.go",
        "service_test.go",
        "third_pendant_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/usersuit/conf",
        "//app/service/main/usersuit/dao/pendant",
        "//app/service/main/usersuit/model",
        "//library/ecode",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "invite.go",
        "medal.go",
        "notify.go",
        "pendant.go",
        "point_flag.go",
        "service.go",
        "third_pendant.go",
    ],
    importpath = "go-common/app/service/main/usersuit/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/coin/api/gorpc",
        "//app/service/main/coin/model",
        "//app/service/main/member/api/gorpc",
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//app/service/main/point/model",
        "//app/service/main/point/rpc/client",
        "//app/service/main/usersuit/conf",
        "//app/service/main/usersuit/dao/invite",
        "//app/service/main/usersuit/dao/medal",
        "//app/service/main/usersuit/dao/pendant",
        "//app/service/main/usersuit/model",
        "//app/service/main/vip/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus",
        "//library/sync/errgroup",
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
