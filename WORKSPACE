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

http_archive(
    name = "rules_pkg",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/1.0.1/rules_pkg-1.0.1.tar.gz",
        "https://github.com/bazelbuild/rules_pkg/releases/download/1.0.1/rules_pkg-1.0.1.tar.gz",
    ],
    sha256 = "d20c951960ed77cb7b341c2a59488534e494d5ad1d30c4818c736d57772a9fef",
)

http_archive(
    name = "rules_oci",
    sha256 = "46ce9edcff4d3d7b3a550774b82396c0fa619cc9ce9da00c1b09a08b45ea5a14",
    strip_prefix = "rules_oci-1.8.0",
    url = "https://github.com/bazel-contrib/rules_oci/releases/download/v1.8.0/rules_oci-v1.8.0.tar.gz",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_rules_dependencies()

go_register_toolchains(version = "1.20.5")

gazelle_dependencies()

load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")

rules_pkg_dependencies()

#################################
# Rules OCI For docker rules
################################

load("@rules_oci//oci:dependencies.bzl", "rules_oci_dependencies")

rules_oci_dependencies()

load("@rules_oci//oci:repositories.bzl", "LATEST_CRANE_VERSION", "oci_register_toolchains")

oci_register_toolchains(
    name = "oci",
    crane_version = LATEST_CRANE_VERSION,
    # Uncommenting the zot toolchain will cause it to be used instead of crane for some tasks.
    # Note that it does not support docker-format images.
    # zot_version = LATEST_ZOT_VERSION,
)

load("@rules_oci//oci:pull.bzl", "oci_pull")

oci_pull(
    name = "distroless_base",
    digest = "sha256:a7af3ef5d69f6534ba0492cc7d6b8fbcffddcb02511b45becc2fac752f907584",
    image = "gcr.io/distroless/base-debian12",
    platforms = [
        "linux/amd64",
    ],
)

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
