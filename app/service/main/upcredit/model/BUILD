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
        "//app/service/main/upcredit/model/calculator:all-srcs",
        "//app/service/main/upcredit/model/canal:all-srcs",
        "//app/service/main/upcredit/model/databus:all-srcs",
        "//app/service/main/upcredit/model/upcrmmodel:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
