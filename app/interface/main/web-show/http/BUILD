package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "http.go",
        "job.go",
        "local.go",
        "notice.go",
        "res.go",
    ],
    importpath = "go-common/app/interface/main/web-show/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/web-show/conf",
        "//app/interface/main/web-show/model/operation",
        "//app/interface/main/web-show/model/resource",
        "//app/interface/main/web-show/service/job",
        "//app/interface/main/web-show/service/operation",
        "//app/interface/main/web-show/service/resource",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/cache",
        "//library/net/http/blademaster/middleware/cache/store",
        "//library/net/http/blademaster/middleware/verify",
        "//library/xstr",
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
