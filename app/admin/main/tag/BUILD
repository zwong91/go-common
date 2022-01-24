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
        "//app/admin/main/tag/cmd:all-srcs",
        "//app/admin/main/tag/conf:all-srcs",
        "//app/admin/main/tag/dao:all-srcs",
        "//app/admin/main/tag/http:all-srcs",
        "//app/admin/main/tag/model:all-srcs",
        "//app/admin/main/tag/service:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
