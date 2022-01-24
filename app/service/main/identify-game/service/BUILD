load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "cloud.go",
        "identify.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/identify-game/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/identify-game/api/grpc/v1:grpc",
        "//app/service/main/identify-game/conf",
        "//app/service/main/identify-game/dao",
        "//app/service/main/identify-game/model",
        "//library/ecode",
        "//library/log",
        "//library/stat",
        "//library/stat/prom",
    ],
)

go_test(
    name = "service_test",
    srcs = [
        "cloud_test.go",
        "identify_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/identify-game/conf",
        "//app/service/main/identify-game/model",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
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
