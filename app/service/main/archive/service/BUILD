package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "archive_test.go",
        "region_pb_test.go",
        "region_test.go",
        "service_test.go",
        "shot_test.go",
        "stat_pb_test.go",
        "upper_pb_test.go",
        "upper_test.go",
        "video_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/archive/conf",
        "//app/service/main/archive/model/archive",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "archive.go",
        "archive_pb.go",
        "region.go",
        "region_pb.go",
        "service.go",
        "shot.go",
        "stat_pb.go",
        "upper.go",
        "upper_pb.go",
        "video.go",
    ],
    importpath = "go-common/app/service/main/archive/service",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/archive/api",
        "//app/service/main/archive/conf",
        "//app/service/main/archive/dao/archive",
        "//app/service/main/archive/dao/share",
        "//app/service/main/archive/dao/videoshot",
        "//app/service/main/archive/model/archive",
        "//library/ecode",
        "//library/log",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//library/time",
        "//vendor/github.com/dgryski/go-farm",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
