package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "charge",
    srcs = [
        "av_charge.go",
        "av_charge_statis.go",
        "bgm.go",
        "column_charge.go",
        "dao.go",
        "date_statia.go",
        "up_charge.go",
    ],
    importpath = "go-common/app/job/main/growup/dao/charge",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/growup/conf",
        "//app/job/main/growup/model/charge",
        "//library/database/sql",
        "//library/log",
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
    name = "charge_test",
    srcs = [
        "av_charge_statis_test.go",
        "av_charge_test.go",
        "bgm_test.go",
        "column_charge_test.go",
        "dao_test.go",
        "date_statia_test.go",
        "up_charge_test.go",
    ],
    embed = [":charge"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/growup/conf",
        "//app/job/main/growup/model/charge",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
