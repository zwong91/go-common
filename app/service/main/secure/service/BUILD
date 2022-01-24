load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "remotelogin.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/secure/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/location/model",
        "//app/service/main/location/rpc/client",
        "//app/service/main/secure/conf",
        "//app/service/main/secure/dao",
        "//app/service/main/secure/model",
        "//library/log",
        "//library/queue/databus",
        "//library/time",
    ],
)

go_test(
    name = "service_test",
    srcs = ["remotelogin_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/secure/conf",
        "//app/service/main/secure/model",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
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
