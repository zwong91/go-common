package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "video_test",
    srcs = [
        "dao_test.go",
        "redis_test.go",
    ],
    embed = [":video"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/favorite/conf",
        "//app/service/main/favorite/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "video",
    srcs = [
        "dao.go",
        "redis.go",
    ],
    importpath = "go-common/app/interface/main/favorite/dao/video",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/favorite/conf",
        "//app/service/main/favorite/model",
        "//library/cache/redis",
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
