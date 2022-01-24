package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "result_test",
    srcs = [
        "archive_test.go",
        "dao_test.go",
        "staff_test.go",
        "video_test.go",
    ],
    embed = [":result"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/archive/conf",
        "//app/job/main/archive/model/archive",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "result",
    srcs = [
        "archive.go",
        "dao.go",
        "staff.go",
        "video.go",
    ],
    importpath = "go-common/app/job/main/archive/dao/result",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/archive/conf",
        "//app/job/main/archive/model/archive",
        "//app/job/main/archive/model/result",
        "//library/database/sql",
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
