package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "pipeline_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/ep/saga/conf",
        "//app/admin/ep/saga/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/github.com/xanzy/go-gitlab",
    ],
)

go_library(
    name = "service",
    srcs = [
        "basic.go",
        "branch.go",
        "collect_project.go",
        "commit.go",
        "config.go",
        "data.go",
        "job.go",
        "member.go",
        "mr.go",
        "pipeline.go",
        "project.go",
        "runner.go",
        "service.go",
        "statistics.go",
        "sync.go",
        "task.go",
        "user.go",
        "wechat.go",
    ],
    importpath = "go-common/app/admin/ep/saga/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/ep/saga/conf",
        "//app/admin/ep/saga/dao",
        "//app/admin/ep/saga/model",
        "//app/admin/ep/saga/service/gitlab",
        "//app/admin/ep/saga/service/utils",
        "//app/admin/ep/saga/service/wechat",
        "//library/cache/memcache",
        "//library/cache/redis",
        "//library/ecode",
        "//library/log",
        "//vendor/github.com/BurntSushi/toml",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/robfig/cron",
        "//vendor/github.com/xanzy/go-gitlab",
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
        "//app/admin/ep/saga/service/gitlab:all-srcs",
        "//app/admin/ep/saga/service/mail:all-srcs",
        "//app/admin/ep/saga/service/utils:all-srcs",
        "//app/admin/ep/saga/service/wechat:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
