package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "agent",
    srcs = [
        "agent.go",
        "dns.go",
        "http.go",
        "ns_endpoint.go",
    ],
    importpath = "go-common/app/service/main/bns/agent",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/bns/agent/backend",
        "//app/service/main/bns/conf",
        "//app/service/main/bns/lib/resolvconf",
        "//app/service/main/bns/lib/shuffle",
        "//library/conf/env",
        "//library/log",
        "//library/stat/prom",
        "//vendor/github.com/miekg/dns",
        "//vendor/github.com/prometheus/client_golang/prometheus/promhttp",
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
        "//app/service/main/bns/agent/backend:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
