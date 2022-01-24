load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

go_library(
    name = "model",
    srcs = [
        "block.go",
        "member.go",
        "monitor.go",
        "official.go",
        "pagination.go",
        "params.go",
        "realname.go",
        "realname_old.go",
        "review.go",
        "search.go",
    ],
    importpath = "go-common/app/admin/main/member/model",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/member/conf",
        "//app/admin/main/member/model/block",
        "//app/service/main/member/model",
        "//app/service/main/member/model/block",
        "//library/log",
        "//library/time",
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
    srcs = [
        ":package-srcs",
        "//app/admin/main/member/model/block:all-srcs",
        "//app/admin/main/member/model/bom:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "model_test",
    srcs = ["member_test.go"],
    embed = [":model"],
    rundir = ".",
    tags = ["automanaged"],
    deps = ["//vendor/github.com/stretchr/testify/assert"],
)
