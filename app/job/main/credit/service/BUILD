package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "blocked_test.go",
        "case_test.go",
        "extra_func_test.go",
        "jury_test.go",
        "reply_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/credit/conf",
        "//app/job/main/credit/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "blocked.go",
        "case.go",
        "credit.go",
        "extra_func.go",
        "jury.go",
        "kpi.go",
        "reply.go",
        "service.go",
    ],
    importpath = "go-common/app/job/main/credit/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/credit/conf",
        "//app/job/main/credit/dao",
        "//app/job/main/credit/model",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/member/api/gorpc",
        "//app/service/main/member/model/block",
        "//library/log",
        "//library/queue/databus",
        "//library/time",
        "//library/xstr",
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
