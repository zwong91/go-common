package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "postprocess",
    srcs = [
        "downGradePostProcessor.go",
        "postprocess.go",
        "relevantInsertPostProcessor.go",
        "scatterTagUpPostProcessor.go",
        "selectInsertPostProcessor.go",
        "weakInterventionPostProcessor.go",
    ],
    importpath = "go-common/app/service/bbq/recsys/service/postprocess",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/bbq/recsys/api/grpc/v1:grpc",
        "//app/service/bbq/recsys/model",
        "//app/service/bbq/recsys/service/retrieve",
        "//app/service/bbq/recsys/service/util",
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
