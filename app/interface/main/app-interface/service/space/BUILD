package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "space_test",
    srcs = [
        "contribute_test.go",
        "service_test.go",
        "space_test.go",
    ],
    embed = [":space"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/model",
        "//app/interface/main/app-interface/model/space",
        "//app/service/main/account/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "space",
    srcs = [
        "contribute.go",
        "service.go",
        "space.go",
    ],
    importpath = "go-common/app/interface/main/app-interface/service/space",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/dao/account",
        "//app/interface/main/app-interface/dao/archive",
        "//app/interface/main/app-interface/dao/article",
        "//app/interface/main/app-interface/dao/audio",
        "//app/interface/main/app-interface/dao/bangumi",
        "//app/interface/main/app-interface/dao/bplus",
        "//app/interface/main/app-interface/dao/coin",
        "//app/interface/main/app-interface/dao/community",
        "//app/interface/main/app-interface/dao/elec",
        "//app/interface/main/app-interface/dao/favorite",
        "//app/interface/main/app-interface/dao/live",
        "//app/interface/main/app-interface/dao/member",
        "//app/interface/main/app-interface/dao/pay",
        "//app/interface/main/app-interface/dao/relation",
        "//app/interface/main/app-interface/dao/search",
        "//app/interface/main/app-interface/dao/shop",
        "//app/interface/main/app-interface/dao/space",
        "//app/interface/main/app-interface/dao/tag",
        "//app/interface/main/app-interface/dao/thumbup",
        "//app/interface/main/app-interface/model",
        "//app/interface/main/app-interface/model/audio",
        "//app/interface/main/app-interface/model/bangumi",
        "//app/interface/main/app-interface/model/bplus",
        "//app/interface/main/app-interface/model/community",
        "//app/interface/main/app-interface/model/favorite",
        "//app/interface/main/app-interface/model/shop",
        "//app/interface/main/app-interface/model/space",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/thumbup/model",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup",
        "//library/time",
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
