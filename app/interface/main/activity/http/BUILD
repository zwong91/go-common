package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "act.go",
        "bnj.go",
        "bws.go",
        "http.go",
        "kfc.go",
        "like.go",
        "match.go",
        "sports.go",
        "timemachine.go",
    ],
    importpath = "go-common/app/interface/main/activity/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/model/bws",
        "//app/interface/main/activity/model/like",
        "//app/interface/main/activity/model/sports",
        "//app/interface/main/activity/service/bnj",
        "//app/interface/main/activity/service/bws",
        "//app/interface/main/activity/service/kfc",
        "//app/interface/main/activity/service/like",
        "//app/interface/main/activity/service/sports",
        "//app/interface/main/activity/service/timemachine",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/metadata",
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
