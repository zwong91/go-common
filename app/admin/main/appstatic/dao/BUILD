load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "redis_test.go",
        "up_nas_test.go",
        "upbfs_test.go",
    ],
    embed = [":dao"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/appstatic/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "redis.go",
        "up_nas.go",
        "upbfs.go",
    ],
    importpath = "go-common/app/admin/main/appstatic/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/appstatic/conf",
        "//app/admin/main/appstatic/model",
        "//library/cache/redis",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
