package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "blacklist.go",
        "bubble.go",
        "cheat.go",
        "combinemails.go",
        "http.go",
        "tag.go",
        "tagincome.go",
        "task_status.go",
        "web.go",
    ],
    importpath = "go-common/app/job/main/growup/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/growup/conf",
        "//app/job/main/growup/service",
        "//app/job/main/growup/service/charge",
        "//app/job/main/growup/service/ctrl",
        "//app/job/main/growup/service/income",
        "//app/job/main/growup/service/tag",
        "//library/log",
        "//library/net/http/blademaster",
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
