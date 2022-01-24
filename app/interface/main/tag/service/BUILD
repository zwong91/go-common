package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "arctag_action_test.go",
        "arctag_changeapi_test.go",
        "arctag_user_test.go",
        "channel_test.go",
        "restag_test.go",
        "service_test.go",
        "sub_sort_test.go",
        "sub_test.go",
        "tag_test.go",
        "tagarc_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/tag/conf",
        "//app/interface/main/tag/model",
        "//app/service/main/archive/api",
        "//library/cache/redis",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "arctag_action.go",
        "arctag_changeapi.go",
        "arctag_check.go",
        "arctag_user.go",
        "channel.go",
        "http.go",
        "new_arc.go",
        "restag.go",
        "rpc.go",
        "rpc_channel.go",
        "rpc_tag.go",
        "service.go",
        "sub.go",
        "sub_sort.go",
        "tag.go",
        "tagarc.go",
    ],
    importpath = "go-common/app/interface/main/tag/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/tag/conf",
        "//app/interface/main/tag/dao",
        "//app/interface/main/tag/model",
        "//app/service/main/account/model",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/figure/model",
        "//app/service/main/figure/rpc/client",
        "//app/service/main/tag/api",
        "//app/service/main/tag/model",
        "//app/service/main/tag/rpc/client",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
        "//library/sync/pipeline/fanout",
        "//library/xstr",
        "//vendor/golang.org/x/sync/singleflight",
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
