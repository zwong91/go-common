package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "bangumi_test",
    srcs = ["bangumi_test.go"],
    embed = [":bangumi"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-channel/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "bangumi",
    srcs = [
        "bangumi.go",
        "grpc.go",
    ],
    importpath = "go-common/app/interface/main/app-channel/dao/bangumi",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-card/model/card/bangumi",
        "//app/interface/main/app-channel/conf",
        "//app/service/openplatform/pgc-season/api/grpc/episode/v1:episode",
        "//app/service/openplatform/pgc-season/api/grpc/season/v1:season",
        "//library/ecode",
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
