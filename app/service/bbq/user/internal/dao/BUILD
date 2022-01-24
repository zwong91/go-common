package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.cache.go",
        "dao.go",
        "forbid_user.go",
        "location.go",
        "monkey.go",
        "user.go",
        "user_base.go",
        "user_black.go",
        "user_card.go",
        "user_fan.go",
        "user_follow.go",
        "user_like.go",
        "user_stat.go",
    ],
    importpath = "go-common/app/service/bbq/user/internal/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/notice-service/api/v1:api",
        "//app/service/bbq/user/api",
        "//app/service/bbq/user/internal/conf",
        "//app/service/bbq/user/internal/model",
        "//app/service/main/account/api",
        "//app/service/main/filter/api/grpc/v1:grpc",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/net/rpc/warden",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/bouk/monkey",
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
        "dao.cache_test.go",
        "dao_test.go",
        "user_base_test.go",
        "user_black_test.go",
        "user_card_test.go",
        "user_fan_test.go",
        "user_follow_test.go",
        "user_like_test.go",
        "user_stat_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/service/bbq/user/api",
        "//app/service/bbq/user/internal/conf",
        "//app/service/bbq/user/internal/model",
        "//library/conf/paladin",
        "//library/log",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
