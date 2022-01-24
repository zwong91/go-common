package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "dao_test",
    srcs = [
        "business_test.go",
        "dao_test.go",
        "mc_test.go",
        "permit_test.go",
        "reason_test.go",
        "tag_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/manager/conf",
        "//app/admin/main/manager/model",
        "//library/ecode",
        "//library/net/http/blademaster/middleware/permit",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "business.go",
        "dao.go",
        "journal.go",
        "mc.go",
        "permit.go",
        "reason.go",
        "tag.go",
    ],
    importpath = "go-common/app/admin/main/manager/dao",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/manager/conf",
        "//app/admin/main/manager/model",
        "//library/cache/memcache",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/permit",
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
