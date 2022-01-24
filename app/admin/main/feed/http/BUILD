load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "bangumi.go",
        "bfs.go",
        "channel.go",
        "common.go",
        "egg.go",
        "http.go",
        "popular.go",
        "search.go",
    ],
    importpath = "go-common/app/admin/main/feed/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/feed/conf",
        "//app/admin/main/feed/model/channel",
        "//app/admin/main/feed/model/common",
        "//app/admin/main/feed/model/egg",
        "//app/admin/main/feed/model/search",
        "//app/admin/main/feed/model/show",
        "//app/admin/main/feed/service/bfs",
        "//app/admin/main/feed/service/channel",
        "//app/admin/main/feed/service/common",
        "//app/admin/main/feed/service/egg",
        "//app/admin/main/feed/service/pgc",
        "//app/admin/main/feed/service/popular",
        "//app/admin/main/feed/service/search",
        "//app/admin/main/feed/util",
        "//app/service/openplatform/pgc-season/api/grpc/episode/v1:episode",
        "//app/service/openplatform/pgc-season/api/grpc/season/v1:season",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
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
