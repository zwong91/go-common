package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "archive.go",
        "cm.go",
        "flow.go",
        "http.go",
        "local.go",
        "monitor.go",
        "music.go",
        "oversea.go",
        "pgc.go",
        "search.go",
        "staff.go",
        "stats.go",
        "task.go",
        "track.go",
        "video.go",
    ],
    importpath = "go-common/app/admin/main/videoup/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/videoup/conf",
        "//app/admin/main/videoup/model/archive",
        "//app/admin/main/videoup/model/monitor",
        "//app/admin/main/videoup/model/music",
        "//app/admin/main/videoup/model/oversea",
        "//app/admin/main/videoup/model/search",
        "//app/admin/main/videoup/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/http/blademaster/render",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/jinzhu/gorm",
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
