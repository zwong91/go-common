package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "cms_test",
    srcs = [
        "dao_test.go",
        "merak_test.go",
        "shelve_test.go",
    ],
    embed = [":cms"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/tv/conf",
        "//library/database/sql",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/h2non/gock.v1:gock_v1",
    ],
)

go_library(
    name = "cms",
    srcs = [
        "dao.go",
        "merak.go",
        "shelve.go",
    ],
    importpath = "go-common/app/job/main/tv/dao/cms",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/tv/conf",
        "//library/database/sql",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/xstr",
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
