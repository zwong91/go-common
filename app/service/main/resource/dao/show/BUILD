package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "show_test",
    srcs = [
        "audit_test.go",
        "card_test.go",
        "dao_test.go",
        "relate_test.go",
        "sidebar_test.go",
    ],
    embed = [":show"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/resource/conf",
        "//library/database/sql",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "show",
    srcs = [
        "audit.go",
        "card.go",
        "dao.go",
        "relate.go",
        "sidebar.go",
    ],
    importpath = "go-common/app/service/main/resource/dao/show",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/resource/conf",
        "//app/service/main/resource/model",
        "//library/database/sql",
        "//library/log",
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
