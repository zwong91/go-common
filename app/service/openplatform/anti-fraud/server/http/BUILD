package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "http_test",
    srcs = [
        "http_test.go",
        "question_test.go",
    ],
    embed = [":http"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/openplatform/anti-fraud/conf",
        "//app/service/openplatform/anti-fraud/model",
        "//library/ecode",
        "//library/net/http/blademaster",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "http",
    srcs = [
        "graph.go",
        "http.go",
        "question.go",
        "risk.go",
    ],
    importpath = "go-common/app/service/openplatform/anti-fraud/server/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/openplatform/anti-fraud/api/grpc/v1:grpc",
        "//app/service/openplatform/anti-fraud/api/http/v1:http",
        "//app/service/openplatform/anti-fraud/conf",
        "//app/service/openplatform/anti-fraud/model",
        "//app/service/openplatform/anti-fraud/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/binding",
        "//library/net/http/blademaster/middleware/verify",
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
