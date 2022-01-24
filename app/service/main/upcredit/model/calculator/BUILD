package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "calculator_test",
    srcs = ["arcfsm_test.go"],
    embed = [":calculator"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/upcredit/conf",
        "//app/service/main/upcredit/model/upcrmmodel",
    ],
)

go_library(
    name = "calculator",
    srcs = [
        "creditscore.go",
        "score_rules.go",
    ],
    importpath = "go-common/app/service/main/upcredit/model/calculator",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/upcredit/common/fsm",
        "//app/service/main/upcredit/conf",
        "//app/service/main/upcredit/mathutil",
        "//app/service/main/upcredit/model/upcrmmodel",
        "//library/log",
        "//library/time",
        "//vendor/github.com/jinzhu/gorm",
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
