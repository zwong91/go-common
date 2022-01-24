package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "upcrm_test",
    srcs = [
        "credit_log_test.go",
        "dao_test.go",
        "scoresection_test.go",
        "sign_up_test.go",
        "task_info_test.go",
        "up_base_info_test.go",
        "up_play_info_test.go",
        "up_rank_test.go",
        "up_scores_history_test.go",
        "up_stats_history_test.go",
    ],
    embed = [":upcrm"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/up/conf",
        "//app/admin/main/up/model/signmodel",
        "//library/net/http/blademaster",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "upcrm",
    srcs = [
        "credit_log.go",
        "dao.go",
        "scoresection.go",
        "sign_up.go",
        "task_info.go",
        "up_base_info.go",
        "up_play_info.go",
        "up_rank.go",
        "up_scores_history.go",
        "up_stats_history.go",
    ],
    importpath = "go-common/app/admin/main/up/dao/upcrm",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/conf",
        "//app/admin/main/up/model/signmodel",
        "//app/admin/main/up/model/upcrmmodel",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/time",
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
