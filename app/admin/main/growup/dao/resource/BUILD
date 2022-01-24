package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "resource",
    srcs = [
        "account.go",
        "article_category.go",
        "resource.go",
        "video_category.go",
        "vip.go",
    ],
    importpath = "go-common/app/admin/main/growup/dao/resource",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/growup/conf",
        "//app/admin/main/growup/model",
        "//app/admin/main/growup/util",
        "//app/service/main/account/api",
        "//app/service/main/vip/model",
        "//app/service/main/vip/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
    name = "resource_test",
    srcs = [
        "account_test.go",
        "article_category_test.go",
        "resource_test.go",
        "video_category_test.go",
        "vip_test.go",
    ],
    embed = [":resource"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/growup/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
