package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "ecode_test",
    srcs = [
        "ecode_test.go",
        "example_test.go",
        "status_test.go",
    ],
    embed = [":ecode"],
    rundir = ".",
    tags = ["manual"],
    deps = [
        "//library/ecode/internal/types",
        "//vendor/github.com/golang/protobuf/ptypes/timestamp",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "ecode",
    srcs = [
        "bbq_ecode.go",
        "common_ecode.go",
        "ecode.go",
        "ep_ecode.go",
        "live_ecode.go",
        "main_ecode.go",
        "open_ecode.go",
        "status.go",
    ],
    importpath = "go-common/library/ecode",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/ecode/internal/types",
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/github.com/golang/protobuf/ptypes",
        "//vendor/github.com/pkg/errors",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["example_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/ecode:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
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
        "//library/ecode/internal/types:all-srcs",
        "//library/ecode/pb:all-srcs",
        "//library/ecode/tip:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
