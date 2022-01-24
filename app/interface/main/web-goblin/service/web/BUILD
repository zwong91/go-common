package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "web_test",
    srcs = [
        "full_test.go",
        "pgc_test.go",
        "recruit_test.go",
        "service_test.go",
        "ugc_test.go",
    ],
    embed = [":web"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/web-goblin/conf",
        "//app/interface/main/web-goblin/model/web",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "web",
    srcs = [
        "channel.go",
        "full.go",
        "pgc.go",
        "recruit.go",
        "service.go",
        "ugc.go",
    ],
    importpath = "go-common/app/interface/main/web-goblin/service/web",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tag/model",
        "//app/interface/main/tag/rpc/client",
        "//app/interface/main/web-goblin/conf",
        "//app/interface/main/web-goblin/dao/web",
        "//app/interface/main/web-goblin/model/web",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
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
