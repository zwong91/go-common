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
        "//app/service/main/bns/lib/resolvconf:all-srcs",
        "//app/service/main/bns/lib/shuffle:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
