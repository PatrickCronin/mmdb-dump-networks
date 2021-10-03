# mmdb-dump-networks

`mmdb-dump-networks` - print every network in an MMDB to STDOUT

![Build](https://github.com/PatrickCronin/mmdb-dump-networks/workflows/Build/badge.svg)
![golangci-lint](https://github.com/PatrickCronin/mmdb-dump-networks/workflows/golangci-lint/badge.svg)
[![go report](https://goreportcard.com/badge/github.com/PatrickCronin/mmdb-dump-networks)](https://goreportcard.com/badge/github.com/PatrickCronin/mmdb-dump-networks)
[![Coverage
Status](https://coveralls.io/repos/github/PatrickCronin/mmdb-dump-networks/badge.svg)](https://coveralls.io/github/PatrickCronin/mmdb-dump-networks)

* [Project Description](#project-description)
* [Usage](#usage)
* [Description](#description)
* [Installation](#installation)
* [Reporting Bugs and Issues](#reporting-bugs-and-issues)
* [Copyright and License](#copyright-and-license)

# Project Description

`mmdb-dump-networks` is a quick-and-dirty command line tool to output the list
of networks stored in an [MMDB](https://github.com/maxmind/MaxMind-DB).

# Usage

## Print Usage

```bash
$ ./mmdb-dump-networks -h
Usage: ./mmdb-dump-networks [-h] [filepath1] [filepath2] ... [filepathN]
```

## Print the List of Networks in an MMDB
```bash
$ ./mmdb-dump-networks GeoIP2-Anonymous-IP-Test.mmdb GeoIP2-Enterprise-Test.mmdb
::27d:a0d8/125
::432b:9c00/120
::4ad1:1000/116
...
```

# Installation

## Binary Releases

Precompiled releases are currently available on our [Releases
page](https://github.com/PatrickCronin/mmdb-dump-networks/releases) for the
following platforms and architectures:

* Linux (i386 and x86_64)
* macOS (x86_64 and arm64)
* Windows (i386 and x86_64)

Look for a release that ends in .tar.gz or .zip. Download the release archive
for your platform and architecture.  Uncompress the archive and you'll see an
eponymous folder. In that folder, you'll find the `mmdb-dump-networks` program.
Copy that program to wherever you want it to live, and start using it.

## Linux Packages

Prebuilt packages are currently available on our [Releases
page](https://github.com/PatrickCronin/mmdb-dump-networks/releases) in the
following formats:

* .deb (Ubuntu or Debian)
* .rpm (RedHat or CentOS)

On Ubuntu or Debian, use `dpkg -i /path/to/the.deb` as root. On RedHat or
CentOS, `rpm -i /path/to/the.rpm` as root. `mmdb-dump-networks` will be
installed in to `/usr/bin/mmdb-dump-networks`.

## Building From Source

`mmdb-dump-networks` is written in Golang, so you'll need a reasonably recent
version of Go (1.16+). This project aims to maintain support for the two most
recent major versions of the Go compiler.

With this in place, simply run:

```bash
$ go install github.com/PatrickCronin/mmdb-dump-networks/cmd/mmdb-dump-networks@latest
```

which will install `mmdb-dump-networks` into the directory named by the `GOBIN`
environment variable, which defaults to `$GOPATH/bin` or `$HOME/go/bin` if the
`GOPATH` environment variable is not set.

# Reporting Bugs and Issues

Bugs and other issues can be reported by filing an issue on our [GitHub issue
tracker](https://github.com/PatrickCronin/routesum/issues).

# Copyright and License

This software is Copyright (c) 2021 by Patrick Cronin.

This is free software, licensed under the terms of the [MIT
License](https://github.com/PatrickCronin/routesum/LICENSE.md).