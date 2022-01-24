package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "cache",
    srcs = [
        "cache.go",
        "tag.go",
        "video_types.go",
    ],
    importpath = "go-common/app/admin/main/up/service/cache",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/conf",
        "//app/interface/main/creative/model/tag",
        "//app/service/main/videoup/model/archive",
        "//library/log",
        "//library/net/http/blademaster",
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
    name = "cache_test",
    srcs = [
        "cache_test.go",
        "tag_test.go",
        "video_types_test.go",
    ],
    embed = [":cache"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/up/conf",
        "//app/interface/main/creative/model/tag",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
