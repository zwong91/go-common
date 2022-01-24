package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "action_test.go",
        "archive_test.go",
        "dm_seg_test.go",
        "dm_sepcial_test.go",
        "dm_test.go",
        "mask_test.go",
        "service_test.go",
        "subject_test.go",
        "subtitle_test.go",
        "task_test.go",
        "track_test.go",
        "transfer_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/dm2/conf",
        "//app/job/main/dm2/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "action.go",
        "archive.go",
        "bnj.go",
        "dm.go",
        "dm_seg.go",
        "dm_special.go",
        "mask.go",
        "service.go",
        "subject.go",
        "subtitle.go",
        "task.go",
        "track.go",
        "transfer.go",
    ],
    importpath = "go-common/app/job/main/dm2/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/dm2/conf",
        "//app/job/main/dm2/dao",
        "//app/job/main/dm2/model",
        "//app/job/main/dm2/model/oplog",
        "//app/service/main/archive/api",
        "//app/service/main/archive/api/gorpc",
        "//app/service/main/archive/model/archive",
        "//app/service/main/filter/api/grpc/v1:grpc",
        "//app/service/main/seq-server/model",
        "//app/service/main/seq-server/rpc/client",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/net/metadata",
        "//library/queue/databus",
        "//library/sync/pipeline/fanout",
        "//library/time",
        "//library/xstr",
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
