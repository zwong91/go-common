package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "upcrmmodel",
    srcs = [
        "creditlog.go",
        "score_section_history.go",
        "up_base_info.go",
        "up_score_history.go",
    ],
    importpath = "go-common/app/service/main/upcredit/model/upcrmmodel",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/time"],
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

go_test(
    name = "upcrmmodel_test",
    srcs = ["creditlog_test.go"],
    embed = [":upcrmmodel"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
