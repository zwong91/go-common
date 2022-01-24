package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "ftp",
    srcs = [
        "dao.go",
        "ftp.go",
    ],
    importpath = "go-common/app/job/main/tv/dao/ftp",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/tv/conf",
        "//library/log",
        "//vendor/github.com/ftp-master",
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
    name = "ftp_test",
    srcs = [
        "dao_test.go",
        "ftp_test.go",
    ],
    embed = [":ftp"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/tv/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
