package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "activity.go",
        "authority.go",
        "auto.go",
        "banner.go",
        "blacklist.go",
        "budgetstatistics.go",
        "charge.go",
        "cheat.go",
        "csv.go",
        "http.go",
        "income.go",
        "local.go",
        "notice.go",
        "offlineactivity.go",
        "special_award.go",
        "tag.go",
        "trade.go",
        "upload.go",
        "web.go",
        "white.go",
        "withdraw.go",
    ],
    importpath = "go-common/app/admin/main/growup/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/growup/conf",
        "//app/admin/main/growup/dao",
        "//app/admin/main/growup/dao/shell",
        "//app/admin/main/growup/model",
        "//app/admin/main/growup/model/offlineactivity",
        "//app/admin/main/growup/service",
        "//app/admin/main/growup/service/income",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/render",
        "//library/time",
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
