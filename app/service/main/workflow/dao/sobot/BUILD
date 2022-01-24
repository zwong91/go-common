package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "sobot_test",
    srcs = [
        "dao_test.go",
        "sobot_test.go",
    ],
    embed = [":sobot"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/workflow/conf",
        "//app/service/main/workflow/model/sobot",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "sobot",
    srcs = [
        "dao.go",
        "sobot.go",
    ],
    importpath = "go-common/app/service/main/workflow/dao/sobot",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/workflow/conf",
        "//app/service/main/workflow/model/sobot",
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
