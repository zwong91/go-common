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
        "audit_result_test.go",
        "intervs_test.go",
        "playurl_test.go",
        "region_test.go",
        "service_test.go",
        "sync_pgc_test.go",
        "uplayurl_test.go",
        "upper_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/tv/conf",
        "//app/admin/main/tv/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "arcType.go",
        "arc_audit.go",
        "archive.go",
        "audit_result.go",
        "full.go",
        "intervs.go",
        "label.go",
        "mango.go",
        "modules.go",
        "order.go",
        "others.go",
        "panel.go",
        "region.go",
        "sear_inter.go",
        "service.go",
        "sync_pgc.go",
        "uplayurl.go",
        "upper.go",
        "user.go",
        "video.go",
        "watermark.go",
    ],
    importpath = "go-common/app/admin/main/tv/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/tv/conf",
        "//app/admin/main/tv/dao",
        "//app/admin/main/tv/model",
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/model/archive",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/siddontang/go-mysql/mysql",
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
