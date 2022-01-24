package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "datadao_test",
    srcs = [
        "api_test.go",
        "cache_test.go",
        "creative_test.go",
        "dao_test.go",
    ],
    embed = [":datadao"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/mcn/conf",
        "//app/interface/main/mcn/dao/global",
        "//app/interface/main/mcn/model/datamodel",
        "//app/interface/main/mcn/tool/datacenter",
        "//library/net/http/blademaster",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "datadao",
    srcs = [
        "api.go",
        "cache.go",
        "creative.go",
        "dao.go",
    ],
    importpath = "go-common/app/interface/main/mcn/dao/datadao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/mcn/conf",
        "//app/interface/main/mcn/dao/cache",
        "//app/interface/main/mcn/dao/global",
        "//app/interface/main/mcn/model/datamodel",
        "//app/interface/main/mcn/model/mcnmodel",
        "//app/interface/main/mcn/tool/cache",
        "//app/interface/main/mcn/tool/datacenter",
        "//app/interface/main/tag/api",
        "//library/cache/memcache",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
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
