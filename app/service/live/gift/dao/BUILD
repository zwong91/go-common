package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "discount_gift_test.go",
        "discount_plan_test.go",
        "gift_online_test.go",
        "gift_plan_test.go",
        "redis_test.go",
        "ugift_day_test.go",
        "ugift_week_test.go",
        "user_gift_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/service/live/gift/api/grpc/v1:grpc",
        "//app/service/live/gift/conf",
        "//app/service/live/gift/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "discount_gift.go",
        "discount_plan.go",
        "gift_online.go",
        "gift_plan.go",
        "redis.go",
        "ugift_day.go",
        "ugift_week.go",
        "user_gift.go",
    ],
    importpath = "go-common/app/service/live/gift/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/activity/api/liverpc",
        "//app/service/live/fans_medal/api/liverpc",
        "//app/service/live/gift/api/grpc/v1:grpc",
        "//app/service/live/gift/conf",
        "//app/service/live/gift/model",
        "//app/service/live/live_user/api/liverpc",
        "//app/service/live/room/api/liverpc",
        "//app/service/live/user/api/liverpc",
        "//app/service/live/xuser/api/grpc/v1:grpc",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/log",
        "//library/net/rpc/liverpc",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
