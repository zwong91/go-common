package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "dao",
    srcs = [
        "dao.go",
        "indexer.go",
        "searcher.go",
    ],
    importpath = "go-common/app/service/main/riot-search/dao",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/riot-search/conf",
        "//app/service/main/riot-search/model",
        "//library/database/sql",
        "//library/log",
        "//vendor/github.com/go-ego/riot",
        "//vendor/github.com/go-ego/riot/types",
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

go_test(
    name = "dao_test",
    srcs = [
        "dao_test.go",
        "indexer_test.go",
        "searcher_test.go",
    ],
    embed = [":dao"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/service/main/riot-search/conf",
        "//app/service/main/riot-search/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
