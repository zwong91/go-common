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
        "//app/admin/main/esports/cmd:all-srcs",
        "//app/admin/main/esports/conf:all-srcs",
        "//app/admin/main/esports/dao:all-srcs",
        "//app/admin/main/esports/http:all-srcs",
        "//app/admin/main/esports/model:all-srcs",
        "//app/admin/main/esports/service:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
