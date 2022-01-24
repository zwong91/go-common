package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "watermark_test",
    srcs = [
        "service_test.go",
        "watermark_test.go",
    ],
    embed = [":watermark"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/model/watermark",
        "//app/interface/main/creative/service",
        "//vendor/github.com/davecgh/go-spew/spew",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "watermark",
    srcs = [
        "service.go",
        "watermark.go",
    ],
    importpath = "go-common/app/interface/main/creative/service/watermark",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//app/interface/main/creative/dao/account",
        "//app/interface/main/creative/dao/bfs",
        "//app/interface/main/creative/dao/drawimg",
        "//app/interface/main/creative/dao/monitor",
        "//app/interface/main/creative/dao/watermark",
        "//app/interface/main/creative/model/newcomer",
        "//app/interface/main/creative/model/watermark",
        "//app/interface/main/creative/service",
        "//app/service/main/account/model",
        "//library/ecode",
        "//library/log",
        "//library/net/metadata",
        "//library/queue/databus",
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
