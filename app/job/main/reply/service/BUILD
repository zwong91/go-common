package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "index_test.go",
        "reply_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/reply/conf",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "admin.go",
        "business.go",
        "fold.go",
        "index.go",
        "notify.go",
        "reply.go",
        "service.go",
        "spam.go",
        "upper.go",
    ],
    importpath = "go-common/app/job/main/reply/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/openplatform/article/model",
        "//app/interface/openplatform/article/rpc/client",
        "//app/job/main/reply/conf",
        "//app/job/main/reply/dao/message",
        "//app/job/main/reply/dao/notice",
        "//app/job/main/reply/dao/reply",
        "//app/job/main/reply/dao/search",
        "//app/job/main/reply/dao/spam",
        "//app/job/main/reply/dao/stat",
        "//app/job/main/reply/model/reply",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/assist/model/assist",
        "//app/service/main/assist/rpc/client",
        "//app/service/main/relation/model",
        "//app/service/openplatform/pgc-season/api/grpc/episode/v1:episode",
        "//library/database/elastic",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/rpc/warden",
        "//library/queue/databus",
        "//library/sync/pipeline/fanout",
        "//library/time",
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
