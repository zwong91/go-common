package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "upcrmdao_test",
    srcs = [
        "dao_test.go",
        "scoresection_test.go",
        "up_base_info_test.go",
    ],
    embed = [":upcrmdao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/upcredit/conf",
        "//app/service/main/upcredit/model/calculator",
        "//vendor/github.com/go-sql-driver/mysql",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "upcrmdao",
    srcs = [
        "dao.go",
        "scoresection.go",
        "up_base_info.go",
    ],
    importpath = "go-common/app/service/main/upcredit/dao/upcrmdao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/upcredit/conf",
        "//app/service/main/upcredit/model/calculator",
        "//app/service/main/upcredit/model/upcrmmodel",
        "//library/log",
        "//library/time",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/siddontang/go-mysql/mysql",
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
