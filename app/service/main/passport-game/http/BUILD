load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "http",
    srcs = [
        "account.go",
        "cache.go",
        "cloud.go",
        "http.go",
        "passport.go",
        "passport_login.go",
        "passport_reg.go",
        "sign.go",
    ],
    importpath = "go-common/app/service/main/passport-game/http",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/service/main/passport-game/conf",
        "//app/service/main/passport-game/model",
        "//app/service/main/passport-game/service",
        "//library/ecode",
        "//library/log",
        "//library/net/http/blademaster",
        "//library/net/metadata",
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
