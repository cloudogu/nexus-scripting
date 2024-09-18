APP=nexus-scripting
VERSION=0.3.0

TARGETDIR=target
PKG=${APP}-${VERSION}.tar.gz
BINARY=${TARGETDIR}/${APP}

default: build

setup:
	dep ensure

$(BINARY): setup
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-X main.Version=${VERSION} -extldflags "-static"' -o $(BINARY) .

build: $(BINARY)

package: build
	cd ${TARGETDIR}; tar cvfz ${PKG} ${APP}

clean:
	rm -rf $(TARGETDIR)
