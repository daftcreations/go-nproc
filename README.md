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

## Build

`go build`
