package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "manager_test",
    srcs = [
        "auth_role_test.go",
        "dao_test.go",
        "httpclient_test.go",
        "manager_test.go",
        "reason_log_test.go",
    ],
    embed = [":manager"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/videoup/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "manager",
    srcs = [
        "auth_role.go",
        "dao.go",
        "httpclient.go",
        "manager.go",
        "reason_log.go",
        "up.go",
    ],
    importpath = "go-common/app/admin/main/videoup/dao/manager",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/videoup/conf",
        "//app/admin/main/videoup/model/manager",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/xstr",
        "//vendor/github.com/pkg/errors",
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
