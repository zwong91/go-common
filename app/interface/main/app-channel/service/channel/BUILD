package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "channel_test",
    srcs = ["channel_test.go"],
    embed = [":channel"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-channel/conf",
        "//app/interface/main/app-channel/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "channel",
    srcs = [
        "audit.go",
        "cache.go",
        "channel.go",
        "index.go",
        "index2.go",
        "infoc.go",
        "manager.go",
        "service.go",
        "tab.go",
    ],
    importpath = "go-common/app/interface/main/app-channel/service/channel",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-card/model",
        "//app/interface/main/app-card/model/card",
        "//app/interface/main/app-card/model/card/audio",
        "//app/interface/main/app-card/model/card/bangumi",
        "//app/interface/main/app-card/model/card/live",
        "//app/interface/main/app-card/model/card/operate",
        "//app/interface/main/app-card/model/card/show",
        "//app/interface/main/app-channel/conf",
        "//app/interface/main/app-channel/dao/account",
        "//app/interface/main/app-channel/dao/activity",
        "//app/interface/main/app-channel/dao/archive",
        "//app/interface/main/app-channel/dao/article",
        "//app/interface/main/app-channel/dao/audio",
        "//app/interface/main/app-channel/dao/audit",
        "//app/interface/main/app-channel/dao/bangumi",
        "//app/interface/main/app-channel/dao/card",
        "//app/interface/main/app-channel/dao/converge",
        "//app/interface/main/app-channel/dao/game",
        "//app/interface/main/app-channel/dao/live",
        "//app/interface/main/app-channel/dao/location",
        "//app/interface/main/app-channel/dao/region",
        "//app/interface/main/app-channel/dao/relation",
        "//app/interface/main/app-channel/dao/shopping",
        "//app/interface/main/app-channel/dao/special",
        "//app/interface/main/app-channel/dao/tab",
        "//app/interface/main/app-channel/dao/tag",
        "//app/interface/main/app-channel/model",
        "//app/interface/main/app-channel/model/activity",
        "//app/interface/main/app-channel/model/card",
        "//app/interface/main/app-channel/model/channel",
        "//app/interface/main/app-channel/model/feed",
        "//app/interface/main/app-channel/model/tab",
        "//app/interface/main/tag/model",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/model/archive",
        "//app/service/main/location/model",
        "//app/service/main/relation/model",
        "//app/service/openplatform/pgc-season/api/grpc/episode/v1:episode",
        "//app/service/openplatform/pgc-season/api/grpc/season/v1:season",
        "//library/log",
        "//library/log/infoc",
        "//library/net/metadata",
        "//library/sync/errgroup",
        "//vendor/github.com/dgryski/go-farm",
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
