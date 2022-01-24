package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = ["service_test.go"],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/tag/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "business.go",
        "res_tag.go",
        "service.go",
        "tag.go",
        "tag_ftp.go",
        "tagarc.go",
    ],
    importpath = "go-common/app/job/main/tag/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/archive/model/archive",
        "//app/job/main/tag/conf",
        "//app/job/main/tag/dao",
        "//app/job/main/tag/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//library/cache",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/xstr",
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
