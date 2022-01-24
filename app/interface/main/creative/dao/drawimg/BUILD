package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "drawimg",
    srcs = [
        "dao.go",
        "draw.go",
        "freetype.go",
        "gaussian.go",
        "imgutil.go",
    ],
    importpath = "go-common/app/interface/main/creative/dao/drawimg",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/interface/main/creative/conf",
        "//library/log",
        "//vendor/github.com/golang/freetype/raster",
        "//vendor/github.com/golang/freetype/truetype",
        "//vendor/golang.org/x/image/font",
        "//vendor/golang.org/x/image/math/fixed",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "drawimg_test",
    srcs = [
        "dao_test.go",
        "draw_test.go",
        "freetype_test.go",
        "imgutil_test.go",
    ],
    embed = [":drawimg"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/interface/main/creative/conf",
        "//vendor/github.com/bouk/monkey",
        "//vendor/github.com/golang/freetype/raster",
        "//vendor/github.com/golang/freetype/truetype",
        "//vendor/github.com/smartystreets/goconvey/convey",
        "//vendor/golang.org/x/image/font",
        "//vendor/golang.org/x/image/math/fixed",
    ],
)
