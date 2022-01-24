package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "account_test",
    srcs = [
        "dao_test.go",
        "memcache_test.go",
        "up_test.go",
    ],
    embed = [":account"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/account",
        "//app/service/main/account/api",
        "//app/service/main/relation/model",
        "//app/service/main/relation/rpc/client",
        "//library/cache/memcache",
        "//library/ecode",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "account",
    srcs = [
        "dao.go",
        "memcache.go",
        "up.go",
    ],
    importpath = "go-common/app/interface/main/creative/dao/account",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/account",
        "//app/interface/main/creative/model/appeal",
        "//app/service/main/account/api",
        "//app/service/main/relation/model",
        "//app/service/main/relation/rpc/client",
        "//library/cache/memcache",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
