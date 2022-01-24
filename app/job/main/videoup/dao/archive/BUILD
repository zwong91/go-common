package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "archive_test",
    srcs = [
        "addit_test.go",
        "archive_biz_test.go",
        "archive_test.go",
        "arhive_history_test.go",
        "config_test.go",
        "dao_test.go",
        "video_test.go",
    ],
    embed = [":archive"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/videoup/conf",
        "//app/job/main/videoup/model/archive",
        "//library/database/sql",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "archive",
    srcs = [
        "addit.go",
        "archive.go",
        "archive_biz.go",
        "archive_history.go",
        "config.go",
        "cover.go",
        "dao.go",
        "databus.go",
        "delay.go",
        "first_pass.go",
        "forbid.go",
        "new_video.go",
        "oper.go",
        "staff.go",
        "stat.go",
        "type.go",
        "video.go",
        "video_audit.go",
        "video_cover.go",
        "video_oper.go",
        "video_shot.go",
    ],
    importpath = "go-common/app/job/main/videoup/dao/archive",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/videoup/conf",
        "//app/job/main/videoup/model/archive",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/xstr",
        "//vendor/github.com/dgryski/go-farm",
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
