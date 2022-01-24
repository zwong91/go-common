package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "dataflow.go",
        "display.go",
        "favorite.go",
        "history.go",
        "http.go",
        "local.go",
        "relation.go",
        "search.go",
        "space.go",
    ],
    importpath = "go-common/app/interface/main/app-interface/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/app-interface/conf",
        "//app/interface/main/app-interface/model",
        "//app/interface/main/app-interface/model/favorite",
        "//app/interface/main/app-interface/model/history",
        "//app/interface/main/app-interface/model/relation",
        "//app/interface/main/app-interface/model/search",
        "//app/interface/main/app-interface/model/space",
        "//app/interface/main/app-interface/service/account",
        "//app/interface/main/app-interface/service/dataflow",
        "//app/interface/main/app-interface/service/display",
        "//app/interface/main/app-interface/service/favorite",
        "//app/interface/main/app-interface/service/history",
        "//app/interface/main/app-interface/service/relation",
        "//app/interface/main/app-interface/service/search",
        "//app/interface/main/app-interface/service/space",
        "//app/interface/main/history/model",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/proxy",
        "//library/net/http/blademaster/middleware/verify",
        "//library/queue/databus",
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
