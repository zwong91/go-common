load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

package(default_visibility = ["//visibility:public"])

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

go_test(
    name = "dao_test",
    srcs = [
        "archive_test.go",
        "dao_test.go",
        "es_test.go",
        "log_test.go",
        "mng_test.go",
        "mng_v2_test.go",
        "query_extra_test.go",
        "query_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/search/conf",
        "//app/admin/main/search/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/gopkg.in/olivere/elastic.v5:elastic_v5",
    ],
)

go_library(
    name = "dao",
    srcs = [
        "archive.go",
        "dao.go",
        "es.go",
        "log.go",
        "mng.go",
        "mng_v2.go",
        "query.go",
        "query_extra.go",
    ],
    importpath = "go-common/app/admin/main/search/dao",
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/search/conf",
        "//app/admin/main/search/model",
        "//library/database/sql",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/stat/prom",
        "//library/sync/errgroup",
        "//vendor/github.com/pkg/errors",
        "//vendor/gopkg.in/olivere/elastic.v5:elastic_v5",
    ],
)
