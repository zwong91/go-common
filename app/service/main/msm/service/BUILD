load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

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
        "app_auth_test.go",
        "config_test.go",
        "ecode_test.go",
        "router_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/infra/config/model",
        "//app/service/main/msm/conf",
        "//app/service/main/msm/model",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "app_auth.go",
        "config.go",
        "ecode.go",
        "router.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/msm/service",
    tags = ["automanaged"],
    deps = [
        "//app/infra/config/model",
        "//app/infra/config/rpc/client",
        "//app/service/main/msm/conf",
        "//app/service/main/msm/dao",
        "//app/service/main/msm/model",
        "//library/ecode",
        "//library/log",
    ],
)
