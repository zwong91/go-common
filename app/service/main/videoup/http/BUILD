package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "archive.go",
        "flow.go",
        "http.go",
        "local.go",
        "netsafe.go",
        "pgc.go",
        "porder.go",
        "reco.go",
        "report.go",
        "staff.go",
        "ugc.go",
        "video.go",
    ],
    importpath = "go-common/app/service/main/videoup/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/videoup/conf",
        "//app/service/main/videoup/model/archive",
        "//app/service/main/videoup/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/antispam",
        "//library/net/http/blademaster/middleware/verify",
        "//library/sync/errgroup",
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
