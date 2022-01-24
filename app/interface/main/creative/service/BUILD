package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "archive_test.go",
        "public_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//library/cache/redis",
        "//vendor/github.com/davecgh/go-spew/spew",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = ["public.go"],
    importpath = "go-common/app/interface/main/creative/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/account",
        "//app/interface/main/creative/dao/activity",
        "//app/interface/main/creative/dao/archive",
        "//app/interface/main/creative/dao/article",
        "//app/interface/main/creative/dao/creative",
        "//app/interface/main/creative/dao/pay",
        "//app/interface/main/creative/dao/subtitle",
        "//app/interface/main/creative/dao/tag",
        "//app/interface/main/creative/dao/up",
        "//app/interface/main/creative/model/activity",
        "//app/interface/main/creative/model/archive",
        "//app/interface/main/creative/model/music",
        "//app/interface/main/creative/model/newcomer",
        "//app/interface/main/creative/model/tag",
        "//app/service/main/account/model",
        "//app/service/main/archive/model/archive",
        "//library/log",
        "//library/queue/databus",
        "//library/sync/errgroup",
        "//library/time",
        "//library/xstr",
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
        "//app/interface/main/creative/service/academy:all-srcs",
        "//app/interface/main/creative/service/account:all-srcs",
        "//app/interface/main/creative/service/ad:all-srcs",
        "//app/interface/main/creative/service/app:all-srcs",
        "//app/interface/main/creative/service/appeal:all-srcs",
        "//app/interface/main/creative/service/archive:all-srcs",
        "//app/interface/main/creative/service/article:all-srcs",
        "//app/interface/main/creative/service/assist:all-srcs",
        "//app/interface/main/creative/service/danmu:all-srcs",
        "//app/interface/main/creative/service/data:all-srcs",
        "//app/interface/main/creative/service/dynamic:all-srcs",
        "//app/interface/main/creative/service/elec:all-srcs",
        "//app/interface/main/creative/service/faq:all-srcs",
        "//app/interface/main/creative/service/feedback:all-srcs",
        "//app/interface/main/creative/service/geetest:all-srcs",
        "//app/interface/main/creative/service/growup:all-srcs",
        "//app/interface/main/creative/service/medal:all-srcs",
        "//app/interface/main/creative/service/music:all-srcs",
        "//app/interface/main/creative/service/newcomer:all-srcs",
        "//app/interface/main/creative/service/operation:all-srcs",
        "//app/interface/main/creative/service/pay:all-srcs",
        "//app/interface/main/creative/service/reply:all-srcs",
        "//app/interface/main/creative/service/resource:all-srcs",
        "//app/interface/main/creative/service/staff:all-srcs",
        "//app/interface/main/creative/service/template:all-srcs",
        "//app/interface/main/creative/service/up:all-srcs",
        "//app/interface/main/creative/service/version:all-srcs",
        "//app/interface/main/creative/service/watermark:all-srcs",
        "//app/interface/main/creative/service/weeklyhonor:all-srcs",
        "//app/interface/main/creative/service/whitelist:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
