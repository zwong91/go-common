package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "notice_test",
    srcs = [
        "audio_playlist_test.go",
        "audio_test.go",
        "bangumi_test.go",
        "credit_test.go",
        "dao_test.go",
        "drawyoo_test.go",
        "dynamic_test.go",
        "live_test.go",
        "topic_test.go",
    ],
    embed = [":notice"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/reply/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "notice",
    srcs = [
        "audio.go",
        "audio_playlist.go",
        "bangumi.go",
        "credit.go",
        "dao.go",
        "drawyoo.go",
        "dynamic.go",
        "live.go",
        "topic.go",
    ],
    importpath = "go-common/app/job/main/reply/dao/notice",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/reply/conf",
        "//library/log",
        "//library/net/http/blademaster",
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
