package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "service",
    srcs = [
        "databus.go",
        "load.go",
        "search.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/riot-search/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/riot-search/conf",
        "//app/service/main/riot-search/dao",
        "//app/service/main/riot-search/model",
        "//library/log",
        "//library/queue/databus",
        "//vendor/github.com/ivpusic/grpool",
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
