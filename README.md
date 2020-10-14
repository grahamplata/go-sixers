[![Go Report Card](https://goreportcard.com/badge/grahamplata/sixers)](https://goreportcard.com/report/grahamplata/sixers)

# Go 76ers

## Current cli

```bash
A cli tool to get past and future statistics about the
Philadelphia 76ers basketball team.
        _____ __
        |___  / /_   ___ _ __ ___
           / / '_ \ / _ \ '__/ __|
          / /| (_) |  __/ |  \__ \
         /_/  \___/ \___|_|  |___/

Usage:
  sixers [command]

Available Commands:
  help        Help about any command
  next        Fetches the next available sixers game date and time.
  record      Fetches the record for the current season.
  schedule    An at a glance view of the Sixers NBA season.

Flags:
      --config string   config file (default is $HOME/.sixers.yaml)
  -h, --help            help for sixers

Use "sixers [command] --help" for more information about a command.
```

## Running tests

```bash
# Test
go test -v ./cmd/...

# Run All tests
go test ./...

# Coverage
go test -cover ./...
```
