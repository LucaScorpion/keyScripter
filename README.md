# keyScripter

A tool to script keyboard and mouse inputs for Windows.

## Usage

To run a script, simply drag it onto the `keyScripter.exe` executable. From the command-line you can run `keyScripter.exe SCRIPT` (where `SCRIPT` is the path to the script). For the full help and options run `keyScripter.exe --help`.

## Scripting Reference

See [SCRIPTING.md](SCRIPTING.md).

## Developing

### Prerequisites

- Make sure you have [Go](https://golang.org/dl/) installed
- Clone the repository

### Building

```shell
go build ./cmd/keyScripter
```

This will create `keyScripter.exe`.

### Formatting

```shell
go fmt ./...
```

### Dependencies

Dependencies are managed using [modules](https://github.com/golang/go/wiki/Modules).
