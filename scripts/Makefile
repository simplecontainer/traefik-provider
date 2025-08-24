.PHONY: test clean

build:
	./development/build.sh

build-docker:
	./development/docker-build.sh

nuke:
	./development/nuke.sh

build: build
docker: nuke build build-docker