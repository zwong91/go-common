package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "award_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/coin/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "award.go",
        "exp.go",
        "service.go",
        "settle.go",
    ],
    importpath = "go-common/app/job/main/coin/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/coin/conf",
        "//app/job/main/coin/dao",
        "//app/job/main/coin/model",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/coin/api/gorpc",
        "//app/service/main/coin/model",
        "//app/service/main/member/api",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/queue/databus/databusutil",
        "//library/stat/prom",
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
