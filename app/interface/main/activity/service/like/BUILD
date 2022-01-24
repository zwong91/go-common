package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "like_test",
    srcs = [
        "act_test.go",
        "like_test.go",
        "match_test.go",
        "missiongroup_test.go",
        "service_test.go",
        "subject_test.go",
    ],
    embed = [":like"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/model/like",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "like",
    srcs = [
        "act.go",
        "bnj.go",
        "like.go",
        "match.go",
        "missiongroup.go",
        "service.go",
        "subject.go",
    ],
    importpath = "go-common/app/interface/main/activity/service/like",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/activity/conf",
        "//app/interface/main/activity/dao/bnj",
        "//app/interface/main/activity/dao/like",
        "//app/interface/main/activity/model/bnj",
        "//app/interface/main/activity/model/like",
        "//app/interface/main/tag/model",
        "//app/interface/main/tag/rpc/client",
        "//app/service/main/account/api",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/coin/api/gorpc",
        "//app/service/main/coin/model",
        "//app/service/main/spy/model",
        "//app/service/main/spy/rpc/client",
        "//app/service/main/thumbup/model",
        "//app/service/main/thumbup/rpc/client",
        "//app/service/main/usersuit/model",
        "//app/service/main/usersuit/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/sync/errgroup",
        "//library/sync/pipeline/fanout",
        "//library/time",
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
