package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "online_filter.go",
        "room_feature.go",
    ],
    importpath = "go-common/app/service/live/recommend/internal/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/recommend/internal/conf",
        "//app/service/live/recommend/recconst",
        "//app/service/live/relation/api/liverpc",
        "//app/service/live/relation/api/liverpc/v1:liverpc",
        "//app/service/live/room/api/liverpc",
        "//app/service/live/room/api/liverpc/v1:liverpc",
        "//app/service/live/room/api/liverpc/v2:liverpc",
        "//library/cache/redis",
        "//library/log",
        "//library/net/rpc/liverpc",
        "//library/sync/errgroup",
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
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "room_feature_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/recommend/internal/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
