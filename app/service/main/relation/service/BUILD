package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "achieve.go",
        "aes.go",
        "audit.go",
        "black.go",
        "follower.go",
        "following.go",
        "infoc.go",
        "monitor.go",
        "prompt.go",
        "relation.go",
        "service.go",
        "stat.go",
        "tag.go",
        "whisper.go",
    ],
    importpath = "go-common/app/service/main/relation/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/member/api/gorpc",
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//app/service/main/relation/conf",
        "//app/service/main/relation/dao",
        "//app/service/main/relation/model",
        "//app/service/main/relation/model/sets",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/net/metadata",
        "//library/stat/prom",
        "//library/xstr",
        "//vendor/github.com/pkg/errors",
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
