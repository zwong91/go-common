package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "timerqueue_test",
    srcs = [
        "example_test.go",
        "queue_test.go",
    ],
    embed = [":timerqueue"],
    rundir = ".",
    tags = ["automanaged"],
)

go_library(
    name = "timerqueue",
    srcs = ["queue.go"],
    importpath = "go-common/app/admin/main/up/util/timerqueue",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = ["//app/admin/main/up/util/timerqueue:go_default_library"],
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
