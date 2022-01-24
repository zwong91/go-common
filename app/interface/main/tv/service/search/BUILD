package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "search",
    srcs = [
        "search.go",
        "service.go",
        "wild.go",
    ],
    importpath = "go-common/app/interface/main/tv/service/search",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/dao/archive",
        "//app/interface/main/tv/dao/cms",
        "//app/interface/main/tv/dao/search",
        "//app/interface/main/tv/model",
        "//app/interface/main/tv/model/search",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup",
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
    name = "search_test",
    srcs = [
        "service_test.go",
        "wild_test.go",
    ],
    embed = [":search"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model/search",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
