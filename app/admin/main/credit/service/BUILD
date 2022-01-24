package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "jury_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/credit/conf",
        "//app/admin/main/credit/model/blocked",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "blocked.go",
        "extra_func.go",
        "jury.go",
        "labour.go",
        "msg.go",
        "service.go",
        "user.go",
    ],
    importpath = "go-common/app/admin/main/credit/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/credit/conf",
        "//app/admin/main/credit/dao/account",
        "//app/admin/main/credit/dao/blocked",
        "//app/admin/main/credit/dao/manager",
        "//app/admin/main/credit/dao/msg",
        "//app/admin/main/credit/dao/relation",
        "//app/admin/main/credit/dao/search",
        "//app/admin/main/credit/dao/upload",
        "//app/admin/main/credit/model",
        "//app/admin/main/credit/model/blocked",
        "//app/service/main/account/api",
        "//app/service/main/account/model",
        "//app/service/main/coin/api",
        "//app/service/main/relation/model",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
        "//library/time",
        "//library/xstr",
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
