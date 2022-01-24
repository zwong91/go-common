package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
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
        "//app/admin/main/workflow/model/manager:all-srcs",
        "//app/admin/main/workflow/model/param:all-srcs",
        "//app/admin/main/workflow/model/search:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
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
        "activity.go",
        "block.go",
        "business.go",
        "business_attr.go",
        "callback.go",
        "challenge.go",
        "common.go",
        "event.go",
        "extra.go",
        "group.go",
        "message.go",
        "platform_meta.go",
        "reply.go",
        "state.go",
        "tag.go",
        "wlog.go",
    ],
    importpath = "go-common/app/admin/main/workflow/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/workflow/model/param",
        "//app/job/main/credit/model",
        "//library/time",
    ],
)
