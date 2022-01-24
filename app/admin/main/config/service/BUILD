package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "app_test.go",
        "build_test.go",
        "common_test.go",
        "conf_svr_test.go",
        "config_test.go",
        "tag_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/config/conf",
        "//app/admin/main/config/model",
        "//vendor/github.com/BurntSushi/toml",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "apm.go",
        "app.go",
        "build.go",
        "common.go",
        "conf_svr.go",
        "config.go",
        "force.go",
        "service.go",
        "tag.go",
        "team.go",
    ],
    importpath = "go-common/app/admin/main/config/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/config/conf",
        "//app/admin/main/config/dao",
        "//app/admin/main/config/model",
        "//app/admin/main/config/pkg/lint",
        "//app/infra/config/model",
        "//app/infra/config/rpc/client",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/satori/go.uuid:go_uuid",
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
