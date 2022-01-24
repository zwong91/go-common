package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "upcrm_test",
    srcs = ["dao_test.go"],
    embed = [":upcrm"],
    tags = ["automanaged"],
    deps = ["//app/job/main/up/conf"],
)

go_library(
    name = "upcrm",
    srcs = [
        "dao.go",
        "sign_up.go",
        "task_info.go",
        "up_base_info.go",
    ],
    importpath = "go-common/app/job/main/up/dao/upcrm",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/up/util/now",
        "//app/job/main/up/conf",
        "//app/job/main/up/model/signmodel",
        "//app/job/main/up/model/upcrmmodel",
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
