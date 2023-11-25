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

Available Commands:
  c           Alias for `copy`
  copy        Copy file to buf dir
  list        List files in buf dir
  p           Alias for `paste`
  paste       Paste files to current dir. Alias: p
  reset       Clear buffered file. If filename is not passed, clear all files in buf dir.

Flags:
      --help      Show help information
      --version   Show version
```
