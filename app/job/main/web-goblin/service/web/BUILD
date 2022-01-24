package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "web",
    srcs = [
        "service.go",
        "ugc.go",
    ],
    importpath = "go-common/app/job/main/web-goblin/service/web",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/web-goblin/conf",
        "//app/job/main/web-goblin/dao/web",
        "//app/job/main/web-goblin/model/web",
        "//app/service/main/dynamic/model",
        "//app/service/main/dynamic/rpc/client",
        "//library/log",
        "//library/queue/databus",
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
    name = "web_test",
    srcs = [
        "service_test.go",
        "ugc_test.go",
    ],
    embed = [":web"],
    tags = ["automanaged"],
    deps = [
        "//app/job/main/web-goblin/conf",
        "//app/job/main/web-goblin/model/web",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
