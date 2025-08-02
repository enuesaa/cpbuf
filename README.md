# cpbuf
A CLI tool to copy and paste files. 

While `cp` would copy and paste in one-time operation, `cpbuf` is stateful.

[![ci](https://github.com/enuesaa/cpbuf/actions/workflows/ci.yaml/badge.svg)](https://github.com/enuesaa/cpbuf/actions/workflows/ci.yaml)


## Install
```bash
go install github.com/enuesaa/cpbuf@v0.0.19
```

## Usage
```console
$ cpbuf --help
A CLI tool to copy and paste files.
`cpbuf` uses a buf dir to hold files temporarily.

Available Commands:
  copy        Copy files to the buf dir (alias: c)
  list        List files in the buf dir (alias: l)
  paste       Paste files into the current dir (alias: p)
  reset       Clear all files in the buf dir (alias: r)

Flags:
  -h, --help      Show help information
  -v, --version   Show version
```

### Copy files
```bash
cpbuf copy a.txt
```

After executing this command, the buf dir `~/.cpbuf` would be created.

### Paste files
```bash
cpbuf paste
```
