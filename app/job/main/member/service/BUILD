package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["service_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/member/conf",
        "//app/job/main/member/model",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "base.go",
        "cache_delay.go",
        "check.go",
        "exp.go",
        "fixer.go",
        "member.go",
        "moral.go",
        "realname.go",
        "service.go",
        "subproc.go",
    ],
    importpath = "go-common/app/job/main/member/service",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/account/service/realname/crypto",
        "//app/job/main/member/conf",
        "//app/job/main/member/dao",
        "//app/job/main/member/model",
        "//app/job/main/member/model/queue",
        "//app/job/main/member/service/block",
        "//app/service/main/member/api/gorpc",
        "//app/service/main/member/model",
        "//app/service/main/share/model",
        "//library/log",
        "//library/log/infoc",
        "//library/net/ip",
        "//library/queue/databus",
        "//library/queue/databus/databusutil",
        "//vendor/github.com/pkg/errors",
        "//vendor/golang.org/x/time/rate",
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
        "//app/job/main/member/service/block:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
