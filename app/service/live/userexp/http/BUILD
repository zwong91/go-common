load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "http",
    srcs = [
        "exp.go",
        "http.go",
    ],
    importpath = "go-common/app/service/live/userexp/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/live/userexp/conf",
        "//app/service/live/userexp/model",
        "//app/service/live/userexp/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/xstr",
    ],
)

go_test(
    name = "http_test",
    srcs = ["http_test.go"],
    embed = [":http"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/live/userexp/conf",
        "//app/service/live/userexp/model",
        "//app/service/live/userexp/service",
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
