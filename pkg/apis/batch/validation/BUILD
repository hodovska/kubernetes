package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = ["validation.go"],
    tags = ["automanaged"],
    deps = [
        "//pkg/api:go_default_library",
        "//pkg/api/validation:go_default_library",
        "//pkg/apis/batch:go_default_library",
        "//pkg/apis/meta/v1:go_default_library",
        "//pkg/apis/meta/v1/validation:go_default_library",
        "//pkg/labels:go_default_library",
        "//pkg/util/validation/field:go_default_library",
        "//vendor:github.com/robfig/cron",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["validation_test.go"],
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//pkg/api:go_default_library",
        "//pkg/apis/batch:go_default_library",
        "//pkg/apis/meta/v1:go_default_library",
        "//pkg/types:go_default_library",
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
)
