# go-nproc

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

- Darwin(Mac) (`arm64`, `x86_64`)
- Linux (`arm64`, `armv6`, `armv7`, `i386`, `x86_64`)
- OpenBSD (`arm64`, `arm6`, `armv7`, `i386`, `x86_64`)
- FreeBSD (`arm64`, `arm6`, `armv7`, `i386`, `x86_64`)
- Windows (`armv6`, `armv7`, `i386`, `x86_64`)
- Solaris (`x86_64`)

```shell
## For linux and mac

# Using wget
wget -O go-nproc https://github.com/pratikbin/go-nproc/releases/download/latest/go-nproc_$(uname -s)_$(uname -m)

# OR

# Using curl
curl -o go-nproc https://github.com/pratikbin/go-nproc/releases/download/latest/go-nproc_$(uname -s)_$(uname -m)

chmod +x ./go-nproc
./go-nproc
```

```shell
## For winodws
# Scrip work in progress, till you can download files from release page as per your architecture
```

## Build

`go build`

## Stargazers over time

[![Stargazers over time](https://starchart.cc/pratikbin/go-nproc.svg)](https://starchart.cc/pratikbin/go-nproc)

---

> > *May the source be with you*
