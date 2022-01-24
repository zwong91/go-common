package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "appeal_test.go",
        "blocked_test.go",
        "jury_test.go",
        "labour_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/credit/conf",
        "//app/interface/main/credit/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "appeal.go",
        "blocked.go",
        "extra_func.go",
        "jury.go",
        "labour.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/credit/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/credit/conf",
        "//app/interface/main/credit/dao",
        "//app/interface/main/credit/model",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/filter/api/grpc/v1:grpc",
        "//app/service/main/member/api/gorpc",
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
        "//library/time",
        "//library/xstr",
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
