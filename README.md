![Continuous Integration](https://github.com/chris-crone/containerized-go-dev/workflows/Continuous%20Integration/badge.svg)

Example Containerized Go Development Environment
------------------------------------------------

This repository contains an example Go project with a containerized development
environment. The example project is a simple CLI tool that echos back its
inputs.

## Why should I containerize my development environment?

There are several advantages to containerizing your development environment:
* You make explicit the tools and versions of tools required to develop your
  project
* Your builds will be more deterministic and reproducible

These will both make it easier for people to collaborate on your project, as
everyone will have the same environment, and make it easier to debug things like
CI failures.

## Prerequisites

The only requirements to build and use this project are Docker and `make`. The
latter can easily be substituted with your scripting tool of choice.

You will also need to enable the BuildKit builder in the Docker CLI. This can be
done by setting `DOCKER_BUILDKIT=1` in your environment.

### macOS

* Install [Docker Desktop](https://www.docker.com/products/docker-desktop)
* Ensure that you have `make` (included with Xcode)
* Run `export DOCKER_BUILDKIT=1` in your terminal or add to your shell
  initialization scripts

### Windows

* Install [Docker Desktop](https://www.docker.com/products/docker-desktop)
* Ensure that you have `make`
* If using PowerShell, run `$env:DOCKER_BUILDKIT=1`
* If using command prompt, run `set DOCKER_BUILDKIT=1`

### Linux

* Install [Docker](https://docs.docker.com/engine/install/)
* Ensure that you have `make`
* Run `export DOCKER_BUILDKIT=1` in your terminal or add to your shell
  initialization scripts

## Getting started

Building the project will output a static binary in the bin/ folder. The
default platform is for macOS but this can be changed using the `PLATFORM` variable:
```console
$ make                        # build for your host OS
$ make PLATFORM=darwin/amd64  # build for macOS
$ make PLATFORM=windows/amd64 # build for Windows x86_64
$ make PLATFORM=linux/amd64   # build for Linux x86_64
$ make PLATFORM=linux/arm     # build for Linux ARM
```

You can then run the binary, which is a simple echo binary, as follows:
```console
$ ./bin/example hello world!
hello world!
```

To run the unit tests run:
```console
$ make unit-test
```

To run the linter:
```console
$ make lint
```

There's then a helpful `test` alias for running both the linter and the unit
tests:
```console
$ make test
```

## Structure of project

### Dockerfile

The [Dockerfile](./Dockerfile) codifies all the tools needed for the project
and the commands that need to be run for building and testing it.

### Makefile

The [Makefile](./Makefile) is purely used to script the required `docker build`
commands as these can get quite long. You can replace this file with a scripting
language of your choice.

### CI

The CI is configured in the [ci.yaml file](./.github/workflows/ci.yaml). By
containerizing the toolchain, the CI relies on the toolchain we defined in the
Dockerfile and doesn't require any custom setup.

## Read more

* [Docker build reference documentation](https://docs.docker.com/engine/reference/commandline/build/)
* [Experimental Dockerfile syntax](https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/experimental.md)
