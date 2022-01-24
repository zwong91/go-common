package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "relation_test",
    srcs = [
        "relation_test.go",
        "service_test.go",
    ],
    embed = [":relation"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/model/relation",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "relation",
    srcs = [
        "relation.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/main/app-interface/service/relation",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/dao/account",
        "//app/interface/main/app-interface/dao/live",
        "//app/interface/main/app-interface/dao/relation",
        "//app/interface/main/app-interface/model/live",
        "//app/interface/main/app-interface/model/relation",
        "//app/service/main/account/model",
        "//app/service/main/relation/model",
        "//library/ecode",
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
