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
        "//app/admin/main/upload/cmd:all-srcs",
        "//app/admin/main/upload/conf:all-srcs",
        "//app/admin/main/upload/dao:all-srcs",
        "//app/admin/main/upload/http:all-srcs",
        "//app/admin/main/upload/model:all-srcs",
        "//app/admin/main/upload/service:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
