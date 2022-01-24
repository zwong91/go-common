package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "callback.go",
        "http.go",
        "push.go",
        "report.go",
        "setting.go",
    ],
    importpath = "go-common/app/interface/main/push/http",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/push/conf",
        "//app/interface/main/push/service",
        "//app/service/main/push/dao/huawei",
        "//app/service/main/push/dao/jpush",
        "//app/service/main/push/dao/mi",
        "//app/service/main/push/dao/oppo",
        "//app/service/main/push/model",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/time",
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
