package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "like.go",
        "relation.go",
        "service.go",
        "user.go",
        "user_face.go",
    ],
    importpath = "go-common/app/service/bbq/user/internal/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/notice-service/api/v1:api",
        "//app/service/bbq/user/api",
        "//app/service/bbq/user/internal/conf",
        "//app/service/bbq/user/internal/dao",
        "//app/service/bbq/user/internal/model",
        "//library/ecode",
        "//library/log",
        "//library/queue/databus",
        "//library/time",
        "//vendor/github.com/golang/protobuf/ptypes/empty",
        "//vendor/github.com/json-iterator/go",
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
    name = "service_test",
    srcs = [
        "like_test.go",
        "relation_test.go",
        "service_test.go",
        "user_test.go",
    ],
    embed = [":service"],
    tags = ["automanaged"],
    deps = [
        "//app/service/bbq/user/api",
        "//app/service/bbq/user/internal/conf",
        "//app/service/bbq/user/internal/model",
        "//library/conf/paladin",
        "//library/ecode",
        "//library/log",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
