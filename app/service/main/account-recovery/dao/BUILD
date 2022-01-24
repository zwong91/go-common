package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "aes.go",
        "captcha.go",
        "dao.go",
        "email.go",
        "mid_info.go",
        "mysql.go",
        "redis.go",
        "req_rpc.go",
        "user_act_log.go",
    ],
    importpath = "go-common/app/service/main/account-recovery/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/account-recovery/conf",
        "//app/service/main/account-recovery/dao/sqlbuilder",
        "//app/service/main/account-recovery/model",
        "//app/service/main/account/api",
        "//app/service/main/location/model",
        "//app/service/main/location/rpc/client",
        "//app/service/main/member/api",
        "//library/cache/redis",
        "//library/database/elastic",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/time",
        "//library/xstr",
        "//vendor/github.com/pkg/errors",
        "//vendor/gopkg.in/gomail.v2:gomail_v2",
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
        "//app/service/main/account-recovery/dao/sqlbuilder:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "dao_test",
    srcs = [
        "captcha_test.go",
        "dao_test.go",
        "email_test.go",
        "mid_info_test.go",
        "mysql_test.go",
        "redis_test.go",
        "req_rpc_test.go",
        "user_act_log_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/account-recovery/conf",
        "//app/service/main/account-recovery/model",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)
