package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "tag_test",
    srcs = ["tag_test.go"],
    embed = [":tag"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-tag/conf",
        "//app/interface/main/app-tag/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "tag",
    srcs = [
        "cache.go",
        "infoc.go",
        "service.go",
        "tag.go",
    ],
    importpath = "go-common/app/interface/main/app-tag/service/tag",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-tag/conf",
        "//app/interface/main/app-tag/dao/archive",
        "//app/interface/main/app-tag/dao/bangumi",
        "//app/interface/main/app-tag/dao/recommend",
        "//app/interface/main/app-tag/dao/region",
        "//app/interface/main/app-tag/dao/tag",
        "//app/interface/main/app-tag/model",
        "//app/interface/main/app-tag/model/bangumi",
        "//app/interface/main/app-tag/model/feed",
        "//app/interface/main/app-tag/model/region",
        "//app/interface/main/app-tag/model/tag",
        "//app/interface/main/tag/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/stat/prom",
        "//library/time",
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
