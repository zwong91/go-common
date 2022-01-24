package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "thirdp_test",
    srcs = [
        "dangbei_test.go",
        "service_test.go",
    ],
    embed = [":thirdp"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/dao/thirdp",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "thirdp",
    srcs = [
        "dangbei.go",
        "mango.go",
        "mango_recom.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/tv/service/thirdp",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/dao/archive",
        "//app/interface/main/tv/dao/cms",
        "//app/interface/main/tv/dao/thirdp",
        "//app/interface/main/tv/model",
        "//app/interface/main/tv/model/thirdp",
        "//app/service/main/archive/api",
        "//library/cache",
        "//library/ecode",
        "//library/log",
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
