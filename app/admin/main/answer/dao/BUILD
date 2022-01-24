load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

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
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "draw.go",
        "es.go",
        "mysql.go",
        "redis.go",
    ],
    importpath = "go-common/app/admin/main/answer/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/answer/conf",
        "//app/admin/main/answer/model",
        "//library/cache/redis",
        "//library/database/elastic",
        "//library/database/orm",
        "//library/log",
        "//library/time",
        "//vendor/github.com/golang/freetype",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/golang.org/x/image/font",
        "//vendor/golang.org/x/image/math/fixed",
    ],
)

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "draw_test.go",
        "es_test.go",
        "mysql_test.go",
        "redis_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/answer/conf",
        "//app/admin/main/answer/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
