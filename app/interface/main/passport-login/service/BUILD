package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "service",
    srcs = [
        "aes.go",
        "cookie.go",
        "cookie_proxy.go",
        "generate.go",
        "md5.go",
        "pwd.go",
        "rsa.go",
        "rsa_time_hash.go",
        "service.go",
        "token.go",
        "token_proxy.go",
        "user.go",
        "user_proxy.go",
        "utils.go",
    ],
    importpath = "go-common/app/interface/main/passport-login/service",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/passport-login/conf",
        "//app/interface/main/passport-login/dao",
        "//app/interface/main/passport-login/model",
        "//library/ecode",
        "//library/log",
        "//vendor/github.com/satori/go.uuid:go_uuid",
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
    name = "service_test",
    srcs = [
        "aes_test.go",
        "cookie_test.go",
        "md5_test.go",
        "pwd_test.go",
        "rsa_test.go",
        "rsa_time_hash_test.go",
        "service_test.go",
        "utils_test.go",
    ],
    embed = [":service"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/passport-login/conf",
        "//app/interface/main/passport-login/model",
        "//vendor/github.com/smartystreets/goconvey/convey",
    ],
)
