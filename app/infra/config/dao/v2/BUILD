package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "app_test.go",
        "build_test.go",
        "config_test.go",
        "dao_test.go",
        "file_test.go",
        "force_test.go",
        "redis_test.go",
        "tag_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/infra/config/conf",
        "//app/infra/config/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "app.go",
        "build.go",
        "config.go",
        "dao.go",
        "file.go",
        "force.go",
        "redis.go",
        "tag.go",
    ],
    importpath = "go-common/app/infra/config/dao/v2",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/infra/config/conf",
        "//app/infra/config/model",
        "//library/cache/redis",
        "//library/database/orm",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/xstr",
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
