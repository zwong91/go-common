package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "upcrmservice_test",
    srcs = [
        "audit_log_test.go",
        "command_test.go",
        "credit_log_test.go",
        "play_info_test.go",
        "scoresection_test.go",
        "sign_up_test.go",
        "up_base_info_test.go",
        "up_rank_test.go",
        "up_stat_history_test.go",
        "upcrmservice_test.go",
        "utils_test.go",
    ],
    embed = [":upcrmservice"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/up/conf",
        "//app/admin/main/up/model",
        "//app/admin/main/up/model/signmodel",
        "//app/admin/main/up/model/upcrmmodel",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "upcrmservice",
    srcs = [
        "audit_log.go",
        "command.go",
        "credit_log.go",
        "play_info.go",
        "scoresection.go",
        "sign_up.go",
        "up_base_info.go",
        "up_rank.go",
        "up_stat_history.go",
        "upcrmservice.go",
        "utils.go",
    ],
    importpath = "go-common/app/admin/main/up/service/upcrmservice",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/conf",
        "//app/admin/main/up/dao/email",
        "//app/admin/main/up/dao/global",
        "//app/admin/main/up/dao/manager",
        "//app/admin/main/up/dao/upcrm",
        "//app/admin/main/up/model",
        "//app/admin/main/up/model/datamodel",
        "//app/admin/main/up/model/signmodel",
        "//app/admin/main/up/model/upcrmmodel",
        "//app/admin/main/up/service/cache",
        "//app/admin/main/up/service/data",
        "//app/admin/main/up/util",
        "//app/admin/main/up/util/mathutil",
        "//app/admin/main/up/util/timerqueue",
        "//app/service/main/account/api",
        "//library/database/elastic",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/queue/databus/report",
        "//library/sync/errgroup",
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
