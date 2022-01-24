package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "newbiedao",
    srcs = [
        "account.go",
        "activity.go",
        "category.go",
        "dao.go",
        "recommend_up.go",
        "relation.go",
        "video_up.go",
    ],
    importpath = "go-common/app/interface/main/growup/dao/newbiedao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/growup/conf",
        "//app/interface/main/growup/model",
        "//app/service/main/account/api",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/xstr",
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
    name = "newbiedao_test",
    srcs = [
        "account_test.go",
        "activity_test.go",
        "category_test.go",
        "dao_test.go",
        "recommend_up_test.go",
        "relation_test.go",
        "video_up_test.go",
    ],
    embed = [":newbiedao"],
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/growup/conf",
        "//app/interface/main/growup/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
