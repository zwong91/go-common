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
        "//app/job/main/tv/model/common:all-srcs",
        "//app/job/main/tv/model/pgc:all-srcs",
        "//app/job/main/tv/model/report:all-srcs",
        "//app/job/main/tv/model/ugc:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
