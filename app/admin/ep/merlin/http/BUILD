package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "audit.go",
        "bilihub.go",
        "cluster.go",
        "dashboard.go",
        "devicefarm.go",
        "http.go",
        "image.go",
        "machine.go",
        "machineV2.go",
        "node.go",
        "tree.go",
        "user.go",
    ],
    importpath = "go-common/app/admin/ep/merlin/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/ep/merlin/conf",
        "//app/admin/ep/merlin/model",
        "//app/admin/ep/merlin/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/middleware/permit",
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
