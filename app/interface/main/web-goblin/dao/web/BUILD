package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "web_test",
    srcs = [
        "channel_test.go",
        "dao_test.go",
        "pgc_test.go",
        "recruit_test.go",
        "ugc_test.go",
    ],
    embed = [":web"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/web-goblin/conf",
        "//app/interface/main/web-goblin/model/web",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "web",
    srcs = [
        "channel.go",
        "dao.go",
        "pgc.go",
        "recruit.go",
        "ugc.go",
    ],
    importpath = "go-common/app/interface/main/web-goblin/dao/web",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/web-goblin/conf",
        "//app/interface/main/web-goblin/model/web",
        "//library/database/elastic",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/stat/prom",
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
