package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "upper_test",
    srcs = [
        "dao_test.go",
        "load_upper_test.go",
        "refresh_test.go",
    ],
    embed = [":upper"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/model/ugc",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "upper",
    srcs = [
        "dao.go",
        "load_upper.go",
        "refresh.go",
        "sync.go",
    ],
    importpath = "go-common/app/job/main/tv/dao/upper",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/model/ugc",
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/stat/prom",
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
