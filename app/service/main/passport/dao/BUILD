package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "hbase_face_test.go",
        "hbase_login_log_test.go",
        "row_key_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/passport/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/github.com/tsuna/gohbase/hrpc",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "hbase_face.go",
        "hbase_login_log.go",
        "hbase_pwd.go",
        "row_key.go",
    ],
    importpath = "go-common/app/service/main/passport/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/passport/conf",
        "//app/service/main/passport/model",
        "//library/database/hbase.v2:hbase_v2",
        "//library/log",
        "//vendor/github.com/tsuna/gohbase/filter",
        "//vendor/github.com/tsuna/gohbase/hrpc",
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
