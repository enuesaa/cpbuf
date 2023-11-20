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
A CLI tool to copy and paste files.
`cpbuf` uses buf-dir to save files temporarily.

Usage:
  cpbuf [flags]
  cpbuf [command]

Available Commands:
  copy        Copy file to buf dir. Alias: c
  list        List files in buf dir
  paste       Paste files to current dir. Alias: p
  reset       Clear buf dir

Flags:
      --help      Show help information
  -v, --version   version for cpbuf
```
