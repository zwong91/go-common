load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "aes.go",
        "auth.go",
        "clean_token.go",
        "contact_bind_log.go",
        "login_log.go",
        "pwd_log.go",
        "service.go",
        "set_token.go",
        "sync_pwd_log.go",
    ],
    importpath = "go-common/app/job/main/passport/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/passport/conf",
        "//app/job/main/passport/dao",
        "//app/job/main/passport/model",
        "//app/service/main/identify-game/model",
        "//app/service/main/identify-game/rpc/client",
        "//library/log",
        "//library/queue/databus",
    ],
)

go_test(
    name = "service_test",
    srcs = [
        "aes_test.go",
        "contact_bind_log_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/passport/conf",
        "//app/job/main/passport/model",
        "//app/service/main/identify-game/model",
        "//library/log",
        "//library/queue/databus",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
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
