load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "account.go",
        "bin_log.go",
        "encrypt_trans.go",
        "passport.go",
        "prom.go",
        "service.go",
    ],
    importpath = "go-common/app/job/main/passport-game-cloud/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/passport-game-cloud/conf",
        "//app/job/main/passport-game-cloud/dao",
        "//app/job/main/passport-game-cloud/model",
        "//library/log",
        "//library/queue/databus",
        "//library/stat/prom",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/pkg/errors",
    ],
)

go_test(
    name = "service_test",
    srcs = [
        "prom_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/passport-game-cloud/conf",
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
