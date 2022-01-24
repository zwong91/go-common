package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "service_test",
    srcs = [
        "ai_test.go",
        "filter_test.go",
        "infoc_test.go",
        "service_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/filter/conf",
        "//app/service/main/filter/model",
        "//app/service/main/filter/model/actriearea",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "service",
    srcs = [
        "admin.go",
        "ai.go",
        "article.go",
        "filter.go",
        "hit.go",
        "infoc.go",
        "key.go",
        "reply.go",
        "rubbish.go",
        "service.go",
    ],
    importpath = "go-common/app/service/main/filter/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/filter/conf",
        "//app/service/main/filter/dao",
        "//app/service/main/filter/model",
        "//app/service/main/filter/model/actriearea",
        "//app/service/main/filter/model/actriekey",
        "//app/service/main/filter/model/lrulist",
        "//app/service/main/filter/model/rpc",
        "//app/service/main/filter/service/ai",
        "//app/service/main/filter/service/area",
        "//app/service/main/filter/service/filter",
        "//app/service/main/filter/service/white",
        "//app/service/main/spy/model",
        "//app/service/main/spy/rpc/client",
        "//library/cache",
        "//library/ecode",
        "//library/log",
        "//library/log/infoc",
        "//library/stat/prom",
        "//vendor/github.com/GRbit/go-pcre",
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
        "//app/service/main/filter/service/ai:all-srcs",
        "//app/service/main/filter/service/area:all-srcs",
        "//app/service/main/filter/service/filter:all-srcs",
        "//app/service/main/filter/service/regexp:all-srcs",
        "//app/service/main/filter/service/white:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
