package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "job_test",
    srcs = ["service_test.go"],
    embed = [":job"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/web-show/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "job",
    srcs = [
        "jobs.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/web-show/service/job",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/web-show/conf",
        "//app/interface/main/web-show/dao/job",
        "//app/interface/main/web-show/model/job",
        "//library/log",
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
