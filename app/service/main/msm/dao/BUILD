load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

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
    name = "dao_test",
    srcs = [
        "app_auth_test.go",
        "dao_test.go",
        "mysql_test.go",
        "tree_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/msm/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "app_auth.go",
        "dao.go",
        "mysql.go",
        "tree.go",
    ],
    importpath = "go-common/app/service/main/msm/dao",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/msm/conf",
        "//app/service/main/msm/model",
        "//library/conf/env",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
    ],
)
