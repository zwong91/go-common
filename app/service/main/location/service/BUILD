package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "ip_old_test.go",
        "ip_test.go",
        "service_test.go",
        "zlimit_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/location/conf",
        "//library/cache/redis",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "anonymous.go",
        "binary.go",
        "ip.go",
        "ip_old.go",
        "service.go",
        "zlimit.go",
    ],
    importpath = "go-common/app/service/main/location/service",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/account/api",
        "//app/service/main/location/conf",
        "//app/service/main/location/dao",
        "//app/service/main/location/model",
        "//library/ecode",
        "//library/log",
        "//library/net/ip",
        "//library/stat/prom",
        "//library/xstr",
        "//vendor/github.com/ipipdotnet/ipdb-go",
        "//vendor/github.com/oschwald/maxminddb-golang",
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
