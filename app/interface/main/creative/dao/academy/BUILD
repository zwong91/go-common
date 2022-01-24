package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "academy_test",
    srcs = [
        "archive_test.go",
        "dao_test.go",
        "feedback_test.go",
        "search_test.go",
        "skill_test.go",
        "tag_test.go",
    ],
    embed = [":academy"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/academy",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "academy",
    srcs = [
        "archive.go",
        "dao.go",
        "feedback.go",
        "search.go",
        "skill.go",
        "tag.go",
    ],
    importpath = "go-common/app/interface/main/creative/dao/academy",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/tool",
        "//app/interface/main/creative/model/academy",
        "//library/database/elastic",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/xstr",
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
