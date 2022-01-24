package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "archive",
    srcs = [
        "archive.go",
        "archive_state.go",
        "biz_poi.go",
        "biz_vote.go",
        "desc_format.go",
        "flow.go",
        "forbid.go",
        "history.go",
        "msg.go",
        "param.go",
        "porder.go",
        "report.go",
        "staff.go",
        "video.go",
    ],
    importpath = "go-common/app/service/main/videoup/model/archive",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/log",
        "//library/time",
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

go_test(
    name = "archive_test",
    srcs = ["video_test.go"],
    embed = [":archive"],
    rundir = ".",
    tags = ["automanaged"],
)
