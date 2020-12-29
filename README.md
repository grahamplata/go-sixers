# Go 76ers

[![Go Report Card](https://goreportcard.com/badge/grahamplata/sixers)](https://goreportcard.com/report/grahamplata/sixers)

A Simple command line tool for my favorite team.

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

## Commands

### Next

```bash
sixers next
10,9 8 76ers! There is a game @ 7:00 PM ET
```

### Record

```bash
sixers record
Wins: 2 Losses: 1 Win pct: 0.667
```

### Schedule

```bash
sixers schedule
---------  ----------------------  -------------  -------------  ------------------------
    Game                    Date           Home           Away                    Winner
---------  ----------------------  -------------  -------------  ------------------------
       1       December 23, 2020       PHI: 113       WAS: 107        Philadelphia 76ers

       2       December 26, 2020        NYK: 89       PHI: 109        Philadelphia 76ers

       3       December 27, 2020       CLE: 118        PHI: 94       Cleveland Cavaliers
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
