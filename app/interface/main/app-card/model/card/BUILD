package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "card",
    srcs = [
        "card.go",
        "double.go",
        "ipad.go",
        "single.go",
    ],
    importpath = "go-common/app/interface/main/app-card/model/card",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-card/model",
        "//app/interface/main/app-card/model/bplus",
        "//app/interface/main/app-card/model/card/ai",
        "//app/interface/main/app-card/model/card/audio",
        "//app/interface/main/app-card/model/card/bangumi",
        "//app/interface/main/app-card/model/card/banner",
        "//app/interface/main/app-card/model/card/cm",
        "//app/interface/main/app-card/model/card/live",
        "//app/interface/main/app-card/model/card/operate",
        "//app/interface/main/app-card/model/card/show",
        "//app/interface/main/tag/model",
        "//app/interface/openplatform/article/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//app/service/main/relation/model",
        "//app/service/openplatform/pgc-season/api/grpc/episode/v1:episode",
        "//app/service/openplatform/pgc-season/api/grpc/season/v1:season",
        "//library/log",
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
    srcs = [
        ":package-srcs",
        "//app/interface/main/app-card/model/card/ai:all-srcs",
        "//app/interface/main/app-card/model/card/audio:all-srcs",
        "//app/interface/main/app-card/model/card/bangumi:all-srcs",
        "//app/interface/main/app-card/model/card/banner:all-srcs",
        "//app/interface/main/app-card/model/card/cm:all-srcs",
        "//app/interface/main/app-card/model/card/live:all-srcs",
        "//app/interface/main/app-card/model/card/operate:all-srcs",
        "//app/interface/main/app-card/model/card/rank:all-srcs",
        "//app/interface/main/app-card/model/card/show:all-srcs",
        "//app/interface/main/app-card/model/card/topic:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
