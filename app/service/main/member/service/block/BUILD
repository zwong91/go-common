package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "block",
    srcs = [
        "audit_log.go",
        "block.go",
        "msg.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/member/service/block",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/member/conf",
        "//app/service/main/member/dao/block",
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//library/database/sql",
        "//library/log",
        "//library/queue/databus/report",
        "//library/sync/pipeline/fanout",
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
    name = "block_test",
    srcs = [
        "block_test.go",
        "service_test.go",
    ],
    embed = [":block"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/member/conf",
        "//app/service/main/member/dao",
        "//app/service/main/member/model/block",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
