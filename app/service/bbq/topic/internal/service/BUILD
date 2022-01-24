package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "extension.go",
        "service.go",
        "topic.go",
    ],
    importpath = "go-common/app/service/bbq/topic/internal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/topic/api",
        "//app/service/bbq/topic/internal/dao",
        "//app/service/bbq/topic/internal/model",
        "//library/conf/paladin",
        "//library/ecode",
        "//library/log",
        "//library/sync/errgroup.v2:errgroup_v2",
        "//vendor/github.com/golang/protobuf/ptypes/empty",
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
    name = "service_test",
    srcs = [
        "extension_test.go",
        "service_test.go",
        "topic_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/service/bbq/topic/api",
        "//app/service/bbq/topic/internal/model",
        "//library/conf/paladin",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
