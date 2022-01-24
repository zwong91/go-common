package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "member.go",
        "passport.go",
        "relation.go",
        "summary.go",
    ],
    importpath = "go-common/app/job/main/account-summary/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/account-summary/conf",
        "//app/job/main/account-summary/model",
        "//app/service/main/member/api/gorpc",
        "//app/service/main/relation/rpc/client",
        "//library/database/hbase.v2:hbase_v2",
        "//library/database/sql",
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
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "member_test.go",
        "passport_test.go",
        "relation_test.go",
        "summary_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/account-summary/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
