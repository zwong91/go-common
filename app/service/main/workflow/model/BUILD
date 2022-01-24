package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "model_test",
    srcs = ["challenge_test.go"],
    embed = [":model"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/smartystreets/goconvey/convey"],
)

go_library(
    name = "model",
    srcs = [
        "archive.go",
        "attachment.go",
        "business.go",
        "callback.go",
        "challenge.go",
        "control.go",
        "event.go",
        "group.go",
        "log.go",
        "tag.go",
    ],
    importpath = "go-common/app/service/main/workflow/model",
    tags = ["automanaged"],
    deps = ["//library/time"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//app/service/main/workflow/model/account:all-srcs",
        "//app/service/main/workflow/model/sobot:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
