package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "warden_test",
    srcs = [
        "exapmle_test.go",
        "logging_test.go",
        "server_test.go",
    ],
    embed = [":warden"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//library/ecode",
        "//library/ecode/pb",
        "//library/ecode/tip",
        "//library/log",
        "//library/net/metadata",
        "//library/net/netutil/breaker",
        "//library/net/rpc/warden/proto/testproto",
        "//library/net/trace",
        "//library/time",
        "//vendor/github.com/golang/protobuf/ptypes",
        "//vendor/github.com/pkg/errors",
        "//vendor/github.com/stretchr/testify/assert",
        "//vendor/google.golang.org/grpc",
        "//vendor/google.golang.org/grpc/codes",
        "//vendor/google.golang.org/grpc/status",
    ],
)

go_library(
    name = "warden",
    srcs = [
        "client.go",
        "logging.go",
        "recovery.go",
        "server.go",
        "stats.go",
        "validate.go",
    ],
    importpath = "go-common/library/net/rpc/warden",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//library/conf/dsn",
        "//library/conf/env",
        "//library/conf/flagvar",
        "//library/ecode",
        "//library/log",
        "//library/naming",
        "//library/naming/discovery",
        "//library/net/metadata",
        "//library/net/netutil/breaker",
        "//library/net/rpc/warden/balancer/wrr",
        "//library/net/rpc/warden/encoding/json",
        "//library/net/rpc/warden/resolver",
        "//library/net/rpc/warden/status",
        "//library/net/trace",
        "//library/stat",
        "//library/stat/summary",
        "//library/stat/sys/cpu",
        "//library/time",
        "//vendor/github.com/pkg/errors",
        "//vendor/google.golang.org/grpc",
        "//vendor/google.golang.org/grpc/codes",
        "//vendor/google.golang.org/grpc/credentials",
        "//vendor/google.golang.org/grpc/keepalive",
        "//vendor/google.golang.org/grpc/metadata",
        "//vendor/google.golang.org/grpc/peer",
        "//vendor/google.golang.org/grpc/reflection",
        "//vendor/google.golang.org/grpc/status",
        "//vendor/gopkg.in/go-playground/validator.v9:validator_v9",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["exapmle_test.go"],
    tags = ["automanaged"],
    deps = [
        "//library/log:go_default_library",
        "//library/net/netutil/breaker:go_default_library",
        "//library/net/rpc/warden:go_default_library",
        "//library/net/rpc/warden/proto/testproto:go_default_library",
        "//library/time:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
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
        "//library/net/rpc/warden/balancer/wrr:all-srcs",
        "//library/net/rpc/warden/benchmark/bench/client:all-srcs",
        "//library/net/rpc/warden/benchmark/bench/proto:all-srcs",
        "//library/net/rpc/warden/benchmark/bench/server:all-srcs",
        "//library/net/rpc/warden/benchmark/helloworld/client:all-srcs",
        "//library/net/rpc/warden/benchmark/helloworld/server:all-srcs",
        "//library/net/rpc/warden/encoding/json:all-srcs",
        "//library/net/rpc/warden/examples/client:all-srcs",
        "//library/net/rpc/warden/examples/grpcDebug:all-srcs",
        "//library/net/rpc/warden/examples/server:all-srcs",
        "//library/net/rpc/warden/metadata:all-srcs",
        "//library/net/rpc/warden/proto/testproto:all-srcs",
        "//library/net/rpc/warden/resolver:all-srcs",
        "//library/net/rpc/warden/status:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
