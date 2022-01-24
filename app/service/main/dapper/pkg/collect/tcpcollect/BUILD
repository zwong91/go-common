package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "tcpcollect_test",
    srcs = ["server_test.go"],
    data = glob(["testdata/**"]),
    embed = [":tcpcollect"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/dapper/conf",
        "//app/service/main/dapper/model",
        "//app/service/main/dapper/pkg/process",
    ],
)

go_library(
    name = "tcpcollect",
    srcs = ["server.go"],
    importpath = "go-common/app/service/main/dapper/pkg/collect/tcpcollect",
    tags = ["manual"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/dapper/conf",
        "//app/service/main/dapper/model",
        "//app/service/main/dapper/pkg/process",
        "//library/log",
        "//library/net/trace/proto",
        "//library/stat/counter",
        "//library/stat/prom",
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/github.com/golang/protobuf/ptypes/duration",
        "//vendor/github.com/golang/protobuf/ptypes/timestamp",
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
