package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "cpm_test",
    srcs = [
        "cpm_test.go",
        "dao_test.go",
    ],
    embed = [":cpm"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/location/model",
        "//app/service/main/resource/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "cpm",
    srcs = [
        "cpm.go",
        "dao.go",
    ],
    importpath = "go-common/app/service/main/resource/dao/cpm",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/location/model",
        "//app/service/main/resource/conf",
        "//app/service/main/resource/model",
        "//library/log",
        "//library/net/http/blademaster",
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
