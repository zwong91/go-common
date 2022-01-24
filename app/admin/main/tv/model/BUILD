load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "model",
    srcs = [
        "archive.go",
        "audit_result.go",
        "channel.go",
        "content_repo.go",
        "full.go",
        "intervs.go",
        "label.go",
        "modules.go",
        "order.go",
        "panel.go",
        "region.go",
        "sear_inter.go",
        "season_repo.go",
        "tv_content.go",
        "uplayurl.go",
        "upper.go",
        "user.go",
        "version_update.go",
        "video.go",
        "water_mark.go",
    ],
    importpath = "go-common/app/admin/main/tv/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/archive/api",
        "//library/database/elastic",
        "//library/time",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/siddontang/go-mysql/mysql",
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
