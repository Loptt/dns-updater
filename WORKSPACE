load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
	name = "io_bazel_rules_go",
	integrity = "sha256-fHbWI2so/2laoozzX5XeMXqUcv0fsUrHl8m/aE8Js3w=",
	urls = [
		"https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.44.2/rules_go-v0.44.2.zip",
		"https://github.com/bazelbuild/rules_go/releases/download/v0.44.2/rules_go-v0.44.2.zip",
	]
)

http_archive(
	name = "bazel_gazelle",
	integrity = "sha256-MpOL2hbmcABjA1R5Bj2dJMYO2o15/Uc5Vj9Q0zHLMgk=",
	urls = [
		"https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
		"https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
	]
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

go_repository(
	name = "com_github_jarcoal_httpmock",
	importpath = "github.com/jarcoal/httpmock",
	sum = "h1:iUx3whfZWVf3jT01hQTO/Eo5sAYtB2/rqaUuOtpInww=",
	version = "v1.3.1"
)

go_repository(
	name = "in_gopkg_yaml_v3",
	importpath = "gopkg.in/yaml.v3",
	sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
	version = "v3.0.1"
)

go_rules_dependencies()

go_register_toolchains(version = "1.20.5")

gazelle_dependencies()

