package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "gorm_test",
    srcs = [
        "business_config_test.go",
        "business_test.go",
        "common_test.go",
        "dao_test.go",
        "direction_test.go",
        "flow_resource_test.go",
        "flow_test.go",
        "net_test.go",
        "report_test.go",
        "resource_test.go",
        "task_config_test.go",
        "task_test.go",
        "token_test.go",
        "transition_test.go",
    ],
    embed = [":gorm"],
    tags = ["automanaged"],
    deps = [
        "//app/admin/main/aegis/conf",
        "//app/admin/main/aegis/model/business",
        "//app/admin/main/aegis/model/common",
        "//app/admin/main/aegis/model/net",
        "//app/admin/main/aegis/model/resource",
        "//app/admin/main/aegis/model/task",
        "//library/ecode",
        "//vendor/github.com/jinzhu/gorm",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)

go_library(
    name = "gorm",
    srcs = [
        "business.go",
        "business_config.go",
        "common.go",
        "dao.go",
        "direction.go",
        "flow.go",
        "flow_resource.go",
        "net.go",
        "report.go",
        "resource.go",
        "task.go",
        "task_config.go",
        "token.go",
        "transition.go",
    ],
    importpath = "go-common/app/admin/main/aegis/dao/gorm",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/admin/main/aegis/conf",
        "//app/admin/main/aegis/model",
        "//app/admin/main/aegis/model/business",
        "//app/admin/main/aegis/model/net",
        "//app/admin/main/aegis/model/resource",
        "//app/admin/main/aegis/model/task",
        "//library/database/orm",
        "//library/ecode",
        "//library/log",
        "//vendor/github.com/jinzhu/gorm",
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
