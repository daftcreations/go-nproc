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
>
> GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
> Report nproc translation bugs to <https://translationproject.org/team/>
> Full documentation at: <https://www.gnu.org/software/coreutils/nproc>
> or available locally via: info '(coreutils) nproc invocation'
> ```

## Install

```shell
## binary
wget -O go-nproc https://github.com/pratikbin/go-nproc/releases/download/0.0.1/go-nproc_0.0.1_$(uname -s)_$(uname -m)
chmod +x ./go-nproc
./go-nproc
```

## Build

`go build`
