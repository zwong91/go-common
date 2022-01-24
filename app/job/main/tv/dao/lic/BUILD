package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "lic",
    srcs = [
        "call.go",
        "dao.go",
        "lic.go",
        "xml.go",
    ],
    importpath = "go-common/app/job/main/tv/dao/lic",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/model/pgc",
        "//library/log",
        "//library/net/http/blademaster",
        "//vendor/github.com/pkg/errors",
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
    name = "lic_test",
    srcs = [
        "dao_test.go",
        "lic_test.go",
        "xml_test.go",
    ],
    embed = [":lic"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/tv/conf",
        "//app/job/main/tv/model/pgc",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
