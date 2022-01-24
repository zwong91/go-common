package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "arc_audit.go",
        "archive.go",
        "audit_result.go",
        "channel.go",
        "content_repo.go",
        "ep_content.go",
        "http.go",
        "intervs.go",
        "label.go",
        "mango.go",
        "modules.go",
        "order.go",
        "panel.go",
        "region.go",
        "sear_inter.go",
        "season_repo.go",
        "upbfs.go",
        "upper.go",
        "user.go",
        "version.go",
        "version_update.go",
        "video.go",
        "watermark.go",
    ],
    importpath = "go-common/app/admin/main/tv/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/tv/conf",
        "//app/admin/main/tv/model",
        "//app/admin/main/tv/service",
        "//library/cache/memcache",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/http/blademaster/render",
        "//library/time",
        "//vendor/github.com/jinzhu/gorm",
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
