# Log Colorizer

**Log Colorizer** is a simple command-line tool written in Go that reads log lines from standard input, detects keywords (such as ERROR, WARNING, INFO, etc.), and outputs colored log lines based on customizable mappings using 256-color ANSI escape codes.


## Default Mappings

If no custom mappings are specified, the tool uses the following defaults:

- **ERROR:** Bright red (`\033[38;5;196m`)
- **WARNING:** Bright yellow (`\033[38;5;226m`)
- **INFO:** Blue (`\033[38;5;33m`)

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) 1.12 or later

### Building

Clone the repository or copy the source code into a file named `main.go`. Then build the binary using:

```bash
go build -o colorize main.go
```

This will produce an executable called `colorize`.

You can also install the bin like so:

```bash
go install
```

## Usage

You can pipe log output to `colorize` to see colored log lines. For example:

```bash
python -u myservice.py | ./colorize
```

### Custom Mappings

Override the default keyword-to-color mappings with the `-mappings` flag. The mappings should be provided as a comma-separated list of `KEY:COLOR` pairs, where `COLOR` is the 256-color code.

For example:

```bash
./colorize -mappings="ERROR:196,WARNING:226,INFO:33,DEBUG:82"
```

In this example:
- `ERROR` messages will be displayed in bright red (color 196).
- `WARNING` messages will be displayed in bright yellow (color 226).
- `INFO` messages will be displayed in blue (color 33).
- `DEBUG` messages will be displayed in green (color 82).
