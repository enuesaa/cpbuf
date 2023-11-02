# cpbuf
A CLI tool to copy and paste files. 

While `cp` would copy and paste in one-time operation, `cpbuf` is stateful.

## Install
```bash
git clone https://github.com/enuesaa/cpbuf.git --depth 1
cd cpbuf
go install
```

## Usage
```console
$ cpbuf --help
A CLI tool to copy and paste files

Usage:
  cpbuf [flags]
  cpbuf [command]

Available Commands:
  clear       clear buf dir
  copy        copy file to buf dir
  list        list filenames in buf dir
  paste       paste files

Flags:
      --help      Show help information
  -v, --version   version for cpbuf
```
