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
        "//app/job/main/appstatic/dao/caldiff:all-srcs",
        "//app/job/main/appstatic/dao/push:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
