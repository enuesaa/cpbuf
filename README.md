# cpbuf
A CLI tool to copy and paste files. 

While `cp` would copy and paste in one-time operation, `cpbuf` is stateful.

[![ci](https://github.com/enuesaa/cpbuf/actions/workflows/ci.yaml/badge.svg)](https://github.com/enuesaa/cpbuf/actions/workflows/ci.yaml)


## Install
```bash
git clone https://github.com/enuesaa/cpbuf.git --depth 1
cd cpbuf
go install
```

## Usage
```console
$ cpbuf --help
A CLI tool to copy and paste files.
`cpbuf` uses buf-dir to save files temporarily.

Available Commands:
  copy        Copy file to buf dir (alias: c)
  list        List files in buf dir (alias: l)
  paste       Paste files to current dir (alias: p)
  reset       Clear buffered files (alias: r)

Flags:
      --help      Show help information
      --version   Show version
```
