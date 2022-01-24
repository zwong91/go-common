package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "show",
    srcs = [
        "app_active.go",
        "card_set_stars.go",
        "channel_tab.go",
        "dao.go",
        "event_topic.go",
        "search_web.go",
        "search_web_card.go",
    ],
    importpath = "go-common/app/admin/main/feed/dao/show",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/feed/conf",
        "//app/admin/main/feed/model/common",
        "//app/admin/main/feed/model/show",
        "//library/database/orm",
        "//library/log",
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

go_test(
    name = "show_test",
    srcs = [
        "app_active_test.go",
        "card_set_stars_test.go",
        "channel_tab_test.go",
        "dao_test.go",
        "event_topic_test.go",
        "search_web_card_test.go",
        "search_web_test.go",
    ],
    embed = [":show"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/feed/conf",
        "//app/admin/main/feed/model/show",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
