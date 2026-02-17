MAKEFILES_VERSION=10.6.0
GOTAG=1.26.0
ARTIFACT_ID=nexus-scripting
VERSION=0.3.0

.DEFAULT_GOAL:=compile-generic

include build/make/variables.mk
include build/make/self-update.mk
include build/make/release.mk
include build/make/prerelease.mk
include build/make/build.mk
include build/make/test-common.mk
include build/make/test-unit.mk
include build/make/package-tar.mk
include build/make/clean.mk
