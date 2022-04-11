# go-nproc

[![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://golang.org/)
[![GitHub license](https://img.shields.io/github/license/daftcreations/go-nproc?style=for-the-badge)](https://github.com/daftcreations/go-nproc/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/v/tag/daftcreations/go-nproc?style=for-the-badge)](https://github.com/daftcreations/go-nproc/releases)
[![Release workflow](https://img.shields.io/github/workflow/status/daftcreation/go-nproc/goreleaser?style=for-the-badge)](https://github.com/daftcreations/go-nproc/actions/workflows/release.yml)
![Lines of code](https://img.shields.io/tokei/lines/github/daftcreations/go-nproc?label=Lines%20of%20code&style=for-the-badge)
[![Chat](https://img.shields.io/discord/960581263264219186?label=%20&logo=discord&style=for-the-badge)](https://discord.com/channels/960581263264219186/960618259244257330)

OS agnostic nproc

> Motivation: Trying macos first time and I can't find `nproc` utility so created one. Tried to create it as similar as [GNU nproc](https://www.gnu.org/software/coreutils/manual/html_node/nproc-invocation.html)
>
> The minimal difference is, this will use single dash `-` for cli flags instead of `--` which GNU uses. I wanted to use standard flags library to keep application simple

## Run

```shell
$ go-nproc -help
go-nproc Usage
  -all
    	print the number of installed processors (default true)
  -help
    	display this help and exit
  -ignore int
    	if possible, exclude N processing units
  -version
    	Get current version
```

> gnu nproc
>
> ```shell
> $ nproc --help
> Usage: nproc [OPTION]...
> Print the number of processing units available to the current process,
> which may be less than the number of online processors
>
>       --all      print the number of installed processors
>       --ignore=N  if possible, exclude N processing units
>       --help     display this help and exit
>       --version  output version information and exit

## Install

Supported platforms

- Darwin(Mac) (`arm64`, `x86_64`) - *tested*
- Linux (`arm64`, `armv6`, `armv7`, `i386`, `x86_64`)  *tested*
- OpenBSD (`arm64`, `arm6`, `armv7`, `i386`, `x86_64`)
- FreeBSD (`arm64`, `arm6`, `armv7`, `i386`, `x86_64`)  *tested*
- Windows (`armv6`, `armv7`, `i386`, `x86_64`)
- Solaris (`x86_64`)

Golang

```shell
go install github.com/daftcreations/go-nproc@latest
```

OR

```shell
## Linux
curl -fsSL instl.sh/daftcreations/go-nproc/linux | bash

## MacOS (m1 and intel)
curl -fsSL instl.sh/daftcreations/go-nproc/macos | bash

## Windows
iwr -useb instl.sh/daftcreations/go-nproc/windows | iex
```

---

[![Discrod](https://img.shields.io/badge/Discord-5865F2?style=for-the-badge&logo=discord&logoColor=white)](https://discord.com/channels/960581263264219186/960618259244257330)
[![Daft Creation](https://img.shields.io/youtube/channel/subscribers/UCDrfHGsm6bJI7Sli7vlcteA?label=YT&logo=youtube&style=for-the-badge)](https://www.youtube.com/c/DaftCreation/playlists)
[![Twitter](https://img.shields.io/twitter/follow/daftcreation?logo=twitter&style=for-the-badge)](https://twitter.com/daftcreation)
[![Instagram](https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white)](https://www.instagram.com/daft.creations/)

### Contributors

![Contributors](https://contrib.rocks/image?repo=daftcreations/go-nproc&columns=80)

### Stargazers over time

[![Stargazers over time](https://starchart.cc/daftcreations/go-nproc.svg)](https://starchart.cc/daftcreations/go-nproc)

---

> > *May the source be with you*
