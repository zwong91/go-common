load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

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
)

go_library(
    name = "syscall",
    srcs = [
        "syscall.go",
        "syscall_windows.go",
    ],
    importpath = "go-common/library/syscall",
    tags = ["automanaged"],
)

go_test(
    name = "syscall_test",
    srcs = [
        "syscall_test.go",
        "syscall_windows_test.go",
    ],
    embed = [":syscall"],
    rundir = ".",
    tags = ["automanaged"],
)
