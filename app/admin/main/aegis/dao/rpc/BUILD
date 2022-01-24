package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "rpc_test",
    srcs = [
        "dao_test.go",
        "task_test.go",
    ],
    embed = [":rpc"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/aegis/conf",
        "//app/service/main/account/api",
        "//app/service/main/relation/model",
        "//app/service/main/up/api/v1:api",
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "rpc",
    srcs = [
        "dao.go",
        "rpc_mock.go",
        "task.go",
    ],
    importpath = "go-common/app/admin/main/aegis/dao/rpc",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/aegis/conf",
        "//app/admin/main/aegis/model",
        "//app/service/main/account/api",
        "//app/service/main/relation/model",
        "//app/service/main/relation/rpc/client",
        "//app/service/main/up/api/v1:api",
        "//library/log",
        "//library/net/metadata",
        "//vendor/github.com/golang/mock/gomock",
        "//vendor/github.com/pkg/errors",
        "//vendor/google.golang.org/grpc",
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
