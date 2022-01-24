package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "pgc",
    srcs = [
        "dao.go",
        "grpc.go",
    ],
    importpath = "go-common/app/admin/main/feed/dao/pgc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/feed/conf",
        "//app/service/openplatform/pgc-season/api/grpc/episode/v1:episode",
        "//app/service/openplatform/pgc-season/api/grpc/season/v1:season",
        "//library/log",
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
    name = "pgc_test",
    srcs = [
        "dao_test.go",
        "grpc_test.go",
    ],
    embed = [":pgc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/feed/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
