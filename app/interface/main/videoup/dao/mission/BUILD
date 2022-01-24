package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "mission_test",
    srcs = [
        "dao_test.go",
        "mission_test.go",
    ],
    embed = [":mission"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/videoup/conf",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "mission",
    srcs = [
        "dao.go",
        "mission.go",
    ],
    importpath = "go-common/app/interface/main/videoup/dao/mission",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/videoup/conf",
        "//app/interface/main/videoup/model/mission",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
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
