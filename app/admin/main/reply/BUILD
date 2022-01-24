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
        "//app/admin/main/reply/cmd:all-srcs",
        "//app/admin/main/reply/conf:all-srcs",
        "//app/admin/main/reply/dao:all-srcs",
        "//app/admin/main/reply/http:all-srcs",
        "//app/admin/main/reply/model:all-srcs",
        "//app/admin/main/reply/service:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
