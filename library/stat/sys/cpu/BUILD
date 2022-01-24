package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "cpu_test",
    srcs = [
        "cgroup_test.go",
        "stat_test.go",
    ],
    embed = [":cpu"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/stretchr/testify/assert"],
)

go_library(
    name = "cpu",
    srcs = [
        "cgroup.go",
        "cpu.go",
        "cpu_darwin.go",
        "cpu_linux.go",
        "cpu_other.go",
        "sysconfig_notcgo.go",
        "util.go",
    ],
    importpath = "go-common/library/stat/sys/cpu",
    tags = ["manual"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//vendor/github.com/pkg/errors",
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
