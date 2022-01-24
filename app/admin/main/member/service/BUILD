load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

go_test(
    name = "service_test",
    srcs = [
        "member_test.go",
        "official_test.go",
        "review_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/member/conf",
        "//app/admin/main/member/model",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "member.go",
        "monitor.go",
        "official.go",
        "realname.go",
        "review.go",
        "review_audit.go",
        "service.go",
    ],
    importpath = "go-common/app/admin/main/member/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/member/conf",
        "//app/admin/main/member/dao",
        "//app/admin/main/member/model",
        "//app/admin/main/member/model/block",
        "//app/admin/main/member/model/bom",
        "//app/admin/main/member/service/block",
        "//app/interface/main/account/service/realname/crypto",
        "//app/service/main/account/api",
        "//app/service/main/coin/api/gorpc",
        "//app/service/main/coin/model",
        "//app/service/main/figure/rpc/client",
        "//app/service/main/member/api/gorpc",
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//app/service/main/member/service/crypto",
        "//app/service/main/relation/model",
        "//app/service/main/relation/rpc/client",
        "//app/service/main/spy/model",
        "//app/service/main/spy/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus",
        "//library/queue/databus/report",
        "//library/stat/prom",
        "//library/time",
        "//vendor/github.com/nfnt/resize",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/robfig/cron",
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
        "//app/admin/main/member/service/block:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
