package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "manager_test",
    srcs = [
        "dao_test.go",
        "manager_test.go",
        "reason_test.go",
        "up_test.go",
    ],
    embed = [":manager"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/videoup/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "manager",
    srcs = [
        "dao.go",
        "manager.go",
        "reason.go",
        "search.go",
        "up.go",
    ],
    importpath = "go-common/app/service/main/videoup/dao/manager",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/videoup/conf",
        "//app/service/main/videoup/model/archive",
        "//library/database/sql",
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
