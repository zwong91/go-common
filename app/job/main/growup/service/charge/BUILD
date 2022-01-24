package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "charge_test",
    srcs = [
        "av_charge_statis_test.go",
        "run_test.go",
        "service_test.go",
    ],
    embed = [":charge"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/growup/conf",
        "//app/job/main/growup/model/charge",
        "//app/job/main/growup/service/ctrl",
        "//library/time",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "charge",
    srcs = [
        "av_charge.go",
        "av_charge_statis.go",
        "bgm_charge.go",
        "column_charge.go",
        "date_statis.go",
        "run.go",
        "run_bgm.go",
        "run_column.go",
        "run_video.go",
        "service.go",
        "trans.go",
        "up_charge.go",
    ],
    importpath = "go-common/app/job/main/growup/service/charge",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/growup/conf",
        "//app/job/main/growup/dao/charge",
        "//app/job/main/growup/dao/email",
        "//app/job/main/growup/model/charge",
        "//app/job/main/growup/service",
        "//app/job/main/growup/service/ctrl",
        "//library/log",
        "//library/time",
        "//vendor/golang.org/x/sync/errgroup",
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
