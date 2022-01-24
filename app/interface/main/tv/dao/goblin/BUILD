package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "goblin",
    srcs = [
        "dao.go",
        "hotword.go",
        "label.go",
        "splash.go",
        "ugc_playurl.go",
        "upgrade.go",
    ],
    importpath = "go-common/app/interface/main/tv/dao/goblin",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model",
        "//app/interface/main/tv/model/goblin",
        "//library/cache/memcache",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
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

go_test(
    name = "goblin_test",
    srcs = [
        "dao_test.go",
        "hotword_test.go",
        "label_test.go",
        "splash_test.go",
        "ugc_playurl_test.go",
        "upgrade_test.go",
    ],
    embed = [":goblin"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tv/conf",
        "//app/interface/main/tv/model",
        "//library/cache/memcache",
        "//library/database/sql",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)
