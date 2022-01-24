package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "archive.go",
        "contest.go",
        "game.go",
        "match.go",
        "match_active.go",
        "match_detail.go",
        "season.go",
        "service.go",
        "tag.go",
        "team.go",
        "tree.go",
    ],
    importpath = "go-common/app/admin/main/esports/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/esports/conf",
        "//app/admin/main/esports/dao",
        "//app/admin/main/esports/model",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
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

go_test(
    name = "service_test",
    srcs = [
        "archive_test.go",
        "game_test.go",
        "match_test.go",
        "service_test.go",
        "team_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/esports/conf",
        "//app/admin/main/esports/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
