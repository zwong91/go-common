package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "business.go",
        "csv.go",
        "debug.go",
        "direction.go",
        "engine.go",
        "flow.go",
        "http.go",
        "monitor.go",
        "net.go",
        "report.go",
        "svg.go",
        "task.go",
        "token.go",
        "transition.go",
    ],
    importpath = "go-common/app/admin/main/aegis/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/aegis/conf",
        "//app/admin/main/aegis/model",
        "//app/admin/main/aegis/model/business",
        "//app/admin/main/aegis/model/common",
        "//app/admin/main/aegis/model/middleware",
        "//app/admin/main/aegis/model/monitor",
        "//app/admin/main/aegis/model/net",
        "//app/admin/main/aegis/model/svg",
        "//app/admin/main/aegis/model/task",
        "//app/admin/main/aegis/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/middleware/permit",
        "//library/net/http/blademaster/middleware/verify",
        "//library/net/http/blademaster/render",
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
