MAKEFILES_VERSION=10.6.0
GOTAG=1.26.0
APP=nexus-scripting
VERSION=0.3.0

.DEFAULT_GOAL:=compile-generic

include build/make/variables.mk
include build/make/self-update.mk
include build/make/release.mk
include build/make/prerelease.mk
include build/make/build.mk
