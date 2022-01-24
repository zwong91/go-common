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
        "//app/interface/main/report-click/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "err_report.go",
        "ip.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/report-click/service",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/history/model",
        "//app/interface/main/history/rpc/client",
        "//app/interface/main/report-click/conf",
        "//app/interface/main/report-click/dao",
        "//app/interface/main/report-click/model",
        "//app/interface/main/report-click/service/crypto/aes",
        "//app/interface/main/report-click/service/crypto/padding",
        "//app/service/main/account/model",
        "//app/service/main/account/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/sync/pipeline/fanout",
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
        "//app/interface/main/report-click/service/crypto/aes:all-srcs",
        "//app/interface/main/report-click/service/crypto/cipher:all-srcs",
        "//app/interface/main/report-click/service/crypto/padding:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
