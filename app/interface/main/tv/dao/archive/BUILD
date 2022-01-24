load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "archive_test",
    srcs = [
        "api_test.go",
        "arctypes_test.go",
        "dao_test.go",
        "memcache_test.go",
        "valid_test.go",
        "view_test.go",
    ],
    embed = [":archive"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model/view",
        "//app/service/main/archive/api",
        "//library/database/sql",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "archive",
    srcs = [
        "api.go",
        "arctypes.go",
        "dao.go",
        "memcache.go",
        "valid.go",
        "view.go",
    ],
    importpath = "go-common/app/interface/main/tv/dao/archive",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/history/model",
        "//app/interface/main/history/rpc/client",
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model",
        "//app/interface/main/tv/model/view",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/xstr",
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
