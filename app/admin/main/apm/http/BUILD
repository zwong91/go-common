package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "http",
    srcs = [
        "app.go",
        "canal.go",
        "dapper.go",
        "databus.go",
        "discovery.go",
        "ecode.go",
        "http.go",
        "monitor.go",
        "need.go",
        "open.go",
        "platform.go",
        "pprof.go",
        "rank.go",
        "upload.go",
        "user.go",
        "ut.go",
        "ut_app.go",
        "ut_dash.go",
    ],
    importpath = "go-common/app/admin/main/apm/http",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/apm/conf",
        "//app/admin/main/apm/model/app",
        "//app/admin/main/apm/model/canal",
        "//app/admin/main/apm/model/databus",
        "//app/admin/main/apm/model/ecode",
        "//app/admin/main/apm/model/need",
        "//app/admin/main/apm/model/pprof",
        "//app/admin/main/apm/model/user",
        "//app/admin/main/apm/model/ut",
        "//app/admin/main/apm/service",
        "//app/tool/saga/model",
        "//library/conf/env",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/middleware/permit",
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

go_test(
    name = "http_test",
    srcs = [
        "apply_test.go",
        "canal_test.go",
        "ecode_test.go",
        "need_test.go",
        "notify_test.go",
        "process_test.go",
        "upload_test.go",
        "user_test.go",
    ],
    embed = [":http"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/apm/conf",
        "//app/admin/main/apm/model/canal",
        "//app/admin/main/apm/model/need",
        "//app/admin/main/apm/service",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
