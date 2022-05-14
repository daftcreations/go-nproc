# go-nproc

<p align="center">
  <a href="https://github.com/daftcreations/go-nproc/blob/master/LICENSE"><img src="https://img.shields.io/github/license/daftcreations/go-nproc?style=flat-square" alt="Licence"/></a>
  <a href="https://github.com/daftcreations/go-nproc/actions/workflows/release.yml"><img src="https://img.shields.io/github/workflow/status/daftcreation/zipper/goreleaser?style=flat-square" alt="build"/></a>
  <a href="https://github.com/daftcreations/go-nproc/releases"><img src="https://img.shields.io/github/v/tag/daftcreations/go-nproc?style=flat-square" alt="tag"/></a>
  <a href="https://goreportcard.com/report/github.com/pratikbalar/zipper"><img src="https://goreportcard.com/badge/github.com/pratikbin/zipper?label=Lines%20of%20code&style=flat-square" alt="go-report"/></a>
  <a href=""><img src="https://img.shields.io/tokei/lines/github/daftcreations/go-nproc?label=Lines%20of%20code&style=flat-square" alt="Lines of code"/></a>
  <a href="https://discord.com/channels/960581263264219186/960618259244257330"><img src="https://img.shields.io/discord/960581263264219186?label=%20&logo=discord&style=flat-square" alt="Discrod"/></a></br
  <a style="text-decoration: none" href="https://github.com/daftcreations/go-nproc/releases"><img src="https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-informational?style=for-the-badge" alt="Downloads"></a>
</p>

OS agnostic nproc (gnu nproc in go)

> size < 400KB

> Motivation: Trying macos first time and I can't find `nproc` utility which will gives you no of logical processor on your system so created one. Tried to create it as similar as [GNU nproc](https://www.gnu.org/software/coreutils/manual/html_node/nproc-invocation.html)
>
> The minimal difference is, this will use single dash `-` for cli flags instead of `--` which GNU uses. I wanted to use standard flags library to keep application simple

## Usage

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

```shell
## Linux
curl -fsSL instl.sh/daftcreations/go-nproc/linux | bash

## MacOS (m1 and intel)
curl -fsSL instl.sh/daftcreations/go-nproc/macos | bash

## Windows
iwr -useb instl.sh/daftcreations/go-nproc/windows | iex
```

Supported platforms, Find binary from release page

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

<div align="center">
    <a href="https://discord.com/channels/960581263264219186/960618259244257330"><img src="https://img.shields.io/badge/Discord-5865F2?style=for-the-badge&logo=discord&logoColor=white"/></a>
    <a href="https://www.youtube.com/c/DaftCreation/playlists"><img src="https://img.shields.io/youtube/channel/subscribers/UCDrfHGsm6bJI7Sli7vlcteA?label=YT&logo=youtube&style=for-the-badge"/></a>
    <a href="https://twitter.com/daftcreations"><img src="https://img.shields.io/twitter/follow/daftcreation?logo=twitter&style=for-the-badge"/></a>
    <a href="https://www.instagram.com/daft.creations/"><img src="https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white"/></a>
</div>

<div align="center">
    <div style="display:flex; justify-content:space-around;">
        <h3 style="margin:-5px 10px 5px;">Contributors</h3>
        <hr align="left" width="20%">
    </div>
    <img src="https://contrib.rocks/image?repo=daftcreations/go-nproc&columns=80" style="width:150px;"/>
</div>

### Stargazers over time

<center>
    <a href="https://starchart.cc/daftcreations/go-nproc"><img src="https://starchart.cc/daftcreations/go-nproc.svg" width="80%"/></a>
</center>

> *May the source be with you*
