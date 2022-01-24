package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "change_history.go",
        "order.go",
        "panel.go",
        "pay.go",
        "qr.go",
        "service.go",
        "token.go",
        "user_info.go",
    ],
    importpath = "go-common/app/service/main/tv/internal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/account/api",
        "//app/service/main/tv/internal/conf",
        "//app/service/main/tv/internal/dao",
        "//app/service/main/tv/internal/model",
        "//app/service/main/tv/internal/service/validator",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/time",
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
        "//app/service/main/tv/internal/service/validator:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "service_test",
    srcs = [
        "change_history_test.go",
        "order_test.go",
        "panel_test.go",
        "pay_test.go",
        "qr_test.go",
        "service_test.go",
        "token_test.go",
        "user_info_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/tv/internal/conf",
        "//app/service/main/tv/internal/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
