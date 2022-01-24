package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "base.go",
        "exp.go",
        "member.go",
        "moral.go",
        "property_review.go",
        "realname.go",
        "service.go",
        "user_flag.go",
    ],
    importpath = "go-common/app/service/main/member/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/member/api",
        "//app/service/main/member/conf",
        "//app/service/main/member/dao",
        "//app/service/main/member/model",
        "//app/service/main/member/service/block",
        "//app/service/main/member/service/crypto",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
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
    srcs = [
        ":package-srcs",
        "//app/service/main/member/service/block:all-srcs",
        "//app/service/main/member/service/crypto:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "service_test",
    srcs = [
        "member_test.go",
        "realname_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/member/conf",
        "//app/service/main/member/dao",
        "//app/service/main/member/model",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
