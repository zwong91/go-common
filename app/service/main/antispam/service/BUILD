package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "keyword_test.go",
        "regexp_content_test.go",
        "service_test.go",
        "transform_test.go",
        "trie_mgr_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/antispam/conf",
        "//app/service/main/antispam/util",
        "//app/service/main/antispam/util/trie",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "service",
    srcs = [
        "cond.go",
        "filter.go",
        "keyword.go",
        "option.go",
        "pipeline.go",
        "regexp.go",
        "rule.go",
        "scheduler.go",
        "service.go",
        "service_impl.go",
        "transform.go",
        "trie_mgr.go",
    ],
    importpath = "go-common/app/service/main/antispam/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/antispam/conf",
        "//app/service/main/antispam/dao",
        "//app/service/main/antispam/extern",
        "//app/service/main/antispam/model",
        "//app/service/main/antispam/util",
        "//app/service/main/antispam/util/trie",
        "//library/cache",
        "//library/log",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "keyword_test.go",
        "regexp_content_test.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor/github.com/stretchr/testify/assert:go_default_library"],
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
