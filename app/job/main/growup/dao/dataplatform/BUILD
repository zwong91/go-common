package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dataplatform_test",
    srcs = [
        "dao_test.go",
        "dataplatform_test.go",
    ],
    embed = [":dataplatform"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/growup/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dataplatform",
    srcs = [
        "dao.go",
        "dataplatform.go",
    ],
    importpath = "go-common/app/job/main/growup/dao/dataplatform",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/growup/conf",
        "//app/job/main/growup/model",
        "//app/job/main/growup/model/income",
        "//library/conf/env",
        "//library/database/sql",
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
