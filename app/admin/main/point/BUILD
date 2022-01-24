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
        "//app/admin/main/point/cmd:all-srcs",
        "//app/admin/main/point/conf:all-srcs",
        "//app/admin/main/point/dao:all-srcs",
        "//app/admin/main/point/http:all-srcs",
        "//app/admin/main/point/model:all-srcs",
        "//app/admin/main/point/service:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
