package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = ["http.go"],
    importpath = "go-common/app/interface/bbq/bullet/internal/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/bbq/bullet/api",
        "//app/interface/bbq/bullet/internal/conf",
        "//app/interface/bbq/bullet/internal/model",
        "//app/interface/bbq/bullet/internal/service",
        "//app/interface/bbq/common/auth",
        "//app/interface/bbq/common/http",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/antispam",
        "//library/net/http/blademaster/middleware/verify",
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
