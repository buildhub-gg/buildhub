load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/buildhub-gg/buildhub/core-be
gazelle(name = "gazelle")

go_library(
    name = "core-be_lib",
    srcs = [
        "server.go",
        "wire_gen.go",
    ],
    importpath = "github.com/buildhub-gg/buildhub/core-be",
    visibility = ["//visibility:private"],
    deps = [
        "//core-be/build",
        "//core-be/graph",
        "//core-be/graph/generated",
        "//core-be/item",
        "@com_github_99designs_gqlgen//graphql/handler:go_default_library",
        "@com_github_99designs_gqlgen//graphql/playground:go_default_library",
        "@com_github_joho_godotenv//autoload:go_default_library",
    ],
)

go_binary(
    name = "core-be",
    embed = [":core-be_lib"],
    visibility = ["//visibility:public"],
)
