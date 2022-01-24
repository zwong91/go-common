package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "audit",
    srcs = [
        "audit.go",
        "dao.go",
        "transcode.go",
    ],
    importpath = "go-common/app/interface/main/tv/dao/audit",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/time",
        "//library/xstr",
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
    name = "audit_test",
    srcs = [
        "audit_test.go",
        "dao_test.go",
        "transcode_test.go",
    ],
    embed = [":audit"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model",
        "//library/database/sql",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
