package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "antispam",
    srcs = ["antispam.go"],
    importpath = "go-common/app/interface/main/upload/http/antispam",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/upload/model",
        "//library/cache/redis",
        "//library/ecode",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
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
    name = "antispam_test",
    srcs = ["antispam_test.go"],
    embed = [":antispam"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/upload/model",
        "//library/cache/redis",
        "//library/container/pool",
        "//library/net/http/blademaster",
        "//library/time",
    ],
)
