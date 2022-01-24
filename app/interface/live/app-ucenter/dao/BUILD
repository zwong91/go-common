package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "fansMedal.go",
        "mainAccount.go",
        "mainFilter.go",
        "relation.go",
        "resource.go",
        "room.go",
        "user.go",
    ],
    importpath = "go-common/app/interface/live/app-ucenter/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/live/app-ucenter/conf",
        "//app/service/live/fans_medal/api/liverpc",
        "//app/service/live/fans_medal/api/liverpc/v1:liverpc",
        "//app/service/live/relation/api/liverpc",
        "//app/service/live/relation/api/liverpc/v1:liverpc",
        "//app/service/live/resource/api/grpc/v1:grpc",
        "//app/service/live/room/api/liverpc",
        "//app/service/live/room/api/liverpc/v1:liverpc",
        "//app/service/live/user/api/liverpc",
        "//app/service/live/user/api/liverpc/v3:liverpc",
        "//app/service/main/member/api",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/net/rpc/liverpc",
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
