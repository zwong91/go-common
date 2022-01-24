package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "feed",
    srcs = [
        "feed.go",
        "player.go",
    ],
    importpath = "go-common/app/interface/main/app-feed/model/feed",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/app-card/model",
        "//app/interface/main/app-card/model/card/ai",
        "//app/interface/main/app-card/model/card/audio",
        "//app/interface/main/app-card/model/card/bangumi",
        "//app/interface/main/app-card/model/card/banner",
        "//app/interface/main/app-card/model/card/cm",
        "//app/interface/main/app-card/model/card/live",
        "//app/interface/main/app-card/model/card/operate",
        "//app/interface/main/app-card/model/card/rank",
        "//app/interface/main/app-card/model/card/show",
        "//app/interface/main/app-feed/model",
        "//app/interface/main/app-feed/model/dislike",
        "//app/interface/main/app-feed/model/live",
        "//app/interface/main/tag/model",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/model/archive",
        "//app/service/main/feed/model",
        "//app/service/main/relation/model",
        "//app/service/openplatform/pgc-season/api/grpc/episode/v1:episode",
        "//library/time",
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
