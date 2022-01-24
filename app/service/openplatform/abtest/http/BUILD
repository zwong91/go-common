package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "http_test",
    srcs = [
        "group_test.go",
        "http_test.go",
        "stra_bench_test.go",
        "stra_test.go",
    ],
    embed = [":http"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/openplatform/abtest/conf",
        "//app/service/openplatform/abtest/service",
        "//library/net/http/blademaster",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "http",
    srcs = [
        "group.go",
        "http.go",
        "stra.go",
    ],
    importpath = "go-common/app/service/openplatform/abtest/http",
    tags = ["automanaged"],
    deps = [
        "//app/service/openplatform/abtest/conf",
        "//app/service/openplatform/abtest/model",
        "//app/service/openplatform/abtest/model/validator",
        "//app/service/openplatform/abtest/service",
        "//library/ecode",
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
