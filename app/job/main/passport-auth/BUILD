package(default_visibility = ["//visibility:public"])

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
        "//app/job/main/passport-auth/cmd:all-srcs",
        "//app/job/main/passport-auth/conf:all-srcs",
        "//app/job/main/passport-auth/dao:all-srcs",
        "//app/job/main/passport-auth/http:all-srcs",
        "//app/job/main/passport-auth/model:all-srcs",
        "//app/job/main/passport-auth/service:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
