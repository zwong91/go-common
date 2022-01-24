package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "question_bank_bind_test.go",
        "question_bank_test.go",
        "question_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/openplatform/anti-fraud/conf",
        "//app/service/openplatform/anti-fraud/model",
        "//library/ecode",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "graph.go",
        "question.go",
        "question_bank.go",
        "question_bank_bind.go",
        "risk.go",
        "service.go",
    ],
    importpath = "go-common/app/service/openplatform/anti-fraud/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/common/openplatform/geetest/model",
        "//app/common/openplatform/geetest/service",
        "//app/common/openplatform/random",
        "//app/service/openplatform/anti-fraud/api/grpc/v1:grpc",
        "//app/service/openplatform/anti-fraud/api/http/v1:http",
        "//app/service/openplatform/anti-fraud/conf",
        "//app/service/openplatform/anti-fraud/dao",
        "//app/service/openplatform/anti-fraud/model",
        "//library/cache/redis",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
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
