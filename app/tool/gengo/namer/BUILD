package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "namer_test",
    srcs = [
        "namer_test.go",
        "plural_namer_test.go",
    ],
    embed = [":namer"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//app/tool/gengo/types"],
)

go_library(
    name = "namer",
    srcs = [
        "doc.go",
        "import_tracker.go",
        "namer.go",
        "order.go",
        "plural_namer.go",
    ],
    importpath = "go-common/app/tool/gengo/namer",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//app/tool/gengo/types"],
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
