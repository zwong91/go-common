package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "alert_test.go",
        "consumer_test.go",
        "monitor_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/openplatform/monitor-end/conf",
        "//app/interface/openplatform/monitor-end/model",
        "//app/interface/openplatform/monitor-end/model/kafka",
        "//app/interface/openplatform/monitor-end/model/monitor",
        "//library/cache/redis",
        "//library/container/pool",
        "//library/database/sql",
        "//library/log/infoc",
        "//library/time",
        "//vendor/github.com/Shopify/sarama",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "alert.go",
        "consumer.go",
        "infoc.go",
        "mail.go",
        "monitor.go",
        "service.go",
    ],
    importpath = "go-common/app/interface/openplatform/monitor-end/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/openplatform/monitor-end/conf",
        "//app/interface/openplatform/monitor-end/dao",
        "//app/interface/openplatform/monitor-end/model",
        "//app/interface/openplatform/monitor-end/model/kafka",
        "//app/interface/openplatform/monitor-end/model/monitor",
        "//app/interface/openplatform/monitor-end/model/prom",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//vendor/github.com/Shopify/sarama",
        "//vendor/github.com/bsm/sarama-cluster",
        "//vendor/github.com/json-iterator/go",
        "//vendor/github.com/scorredoira/email",
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
