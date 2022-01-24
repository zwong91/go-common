package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "http",
    srcs = [
        "cond.go",
        "filter.go",
        "http.go",
        "keyword.go",
        "protocol.go",
        "regexp.go",
        "rule.go",
    ],
    importpath = "go-common/app/service/main/antispam/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/antispam/conf",
        "//app/service/main/antispam/dao",
        "//app/service/main/antispam/model",
        "//app/service/main/antispam/service",
        "//app/service/main/antispam/util",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/middleware/auth",
        "//library/net/http/blademaster/middleware/verify",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["cond_test.go"],
    tags = ["automanaged"],
    deps = [
        "//app/service/main/antispam/http:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
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

go_test(
    name = "http_test",
    srcs = ["cond_test.go"],
    deps = [
        ":http",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)
