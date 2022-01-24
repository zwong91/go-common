package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "bangumi_test",
    srcs = ["dao_test.go"],
    embed = [":bangumi"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "bangumi",
    srcs = [
        "dao.go",
        "grpc.go",
    ],
    importpath = "go-common/app/interface/main/app-interface/dao/bangumi",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/model/bangumi",
        "//app/service/openplatform/pgc-season/api/grpc/season/v1:season",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
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
