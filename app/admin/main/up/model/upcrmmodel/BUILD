package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "upcrmmodel",
    srcs = [
        "creditlog.go",
        "requestargs.go",
        "score_section_history.go",
        "task_info.go",
        "up_base_info.go",
        "up_play_info.go",
        "up_rank.go",
        "up_scores_history.go",
        "up_stats_history.go",
    ],
    importpath = "go-common/app/admin/main/up/model/upcrmmodel",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/model/datamodel",
        "//app/admin/main/up/util",
        "//library/time",
        "//vendor/github.com/siddontang/go-mysql/mysql",
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
