package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "balancer.go",
        "conn.go",
        "online.go",
        "push.go",
        "server.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/broadcast/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/broadcast/api/grpc/v1:grpc",
        "//app/service/main/broadcast/dao",
        "//app/service/main/broadcast/model",
        "//app/service/main/identify/api/grpc",
        "//app/service/main/location/api",
        "//library/cache",
        "//library/conf/paladin",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/naming",
        "//library/naming/discovery",
        "//library/net/metadata",
        "//library/net/netutil",
        "//library/net/rpc/warden",
        "//library/net/rpc/warden/balancer/wrr",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/zhenjl/cityhash",
        "//vendor/google.golang.org/grpc",
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
    name = "service_test",
    srcs = [
        "balancer_test.go",
        "conn_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/broadcast/model",
        "//library/conf/paladin",
        "//library/naming/discovery",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
