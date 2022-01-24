package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "model",
    srcs = [
        "engine.go",
        "log.go",
        "report.go",
    ],
    importpath = "go-common/app/admin/main/aegis/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/aegis/model/common",
        "//app/admin/main/aegis/model/net",
        "//app/admin/main/aegis/model/resource",
        "//app/admin/main/aegis/model/task",
        "//app/service/main/account/api",
        "//library/log",
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
    srcs = [
        ":package-srcs",
        "//app/admin/main/aegis/model/business:all-srcs",
        "//app/admin/main/aegis/model/common:all-srcs",
        "//app/admin/main/aegis/model/databus:all-srcs",
        "//app/admin/main/aegis/model/middleware:all-srcs",
        "//app/admin/main/aegis/model/monitor:all-srcs",
        "//app/admin/main/aegis/model/net:all-srcs",
        "//app/admin/main/aegis/model/resource:all-srcs",
        "//app/admin/main/aegis/model/svg:all-srcs",
        "//app/admin/main/aegis/model/task:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "model_test",
    srcs = ["model_test.go"],
    embed = [":model"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/aegis/model/common",
        "//app/admin/main/aegis/model/resource",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
