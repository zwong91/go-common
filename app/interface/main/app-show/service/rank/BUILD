package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "rank_test",
    srcs = ["rank_test.go"],
    embed = [":rank"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-show/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "rank",
    srcs = [
        "audit.go",
        "rank.go",
    ],
    importpath = "go-common/app/interface/main/app-show/service/rank",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-show/conf",
        "//app/interface/main/app-show/dao/account",
        "//app/interface/main/app-show/dao/archive",
        "//app/interface/main/app-show/dao/audit",
        "//app/interface/main/app-show/dao/recommend",
        "//app/interface/main/app-show/dao/region",
        "//app/interface/main/app-show/dao/relation",
        "//app/interface/main/app-show/model",
        "//app/interface/main/app-show/model/region",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//app/service/main/relation/model",
        "//library/log",
        "//library/sync/errgroup",
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
