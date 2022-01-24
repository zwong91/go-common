package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "server",
    srcs = [
        "bucket.go",
        "channel.go",
        "errors.go",
        "noauth.go",
        "operation.go",
        "report.go",
        "ring.go",
        "room.go",
        "round.go",
        "server.go",
        "server_tcp.go",
        "server_tcp_v1.go",
        "server_websocket.go",
        "server_websocket_v1.go",
        "whitelist.go",
    ],
    importpath = "go-common/app/interface/main/broadcast/server",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/broadcast/conf",
        "//app/interface/main/broadcast/model",
        "//app/service/main/broadcast/api/grpc/v1:grpc",
        "//app/service/main/broadcast/libs/bufio",
        "//app/service/main/broadcast/libs/bytes",
        "//app/service/main/broadcast/libs/time",
        "//app/service/main/broadcast/libs/websocket",
        "//app/service/main/broadcast/model",
        "//library/conf/env",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/http/blademaster/render",
        "//library/net/metadata",
        "//library/net/netutil",
        "//vendor/github.com/gogo/protobuf/proto",
        "//vendor/github.com/gogo/protobuf/types",
        "//vendor/github.com/google/uuid",
        "//vendor/github.com/zhenjl/cityhash",
        "//vendor/google.golang.org/grpc",
        "//vendor/google.golang.org/grpc/encoding/gzip",
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
        "//app/interface/main/broadcast/server/grpc:all-srcs",
        "//app/interface/main/broadcast/server/http:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "server_test",
    srcs = ["report_test.go"],
    embed = [":server"],
    tags = ["automanaged"],
)
