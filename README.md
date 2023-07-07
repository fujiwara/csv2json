# csv2json

## Overview

csv2json is a simple command line interface (CLI) tool that converts CSV files to JSON format.

## Installation

csv2json can be installed using the `go install` command.

```bash
$ go install github.com/fujiwara/csv2json@latest
```

## How to Use

csv2json can be used either by specifying a CSV file as an argument or by piping data in through standard input.

```bash
$ csv2json file.csv
```

```bash
$ cat file.csv | csv2json
```

The tool reads the CSV data, parses it, and then outputs the data in JSON format to standard output.

## Flags

`-n`
The -n flag specifies that the CSV file does not have a header row. If this flag is set, each row is output as a simple JSON array.

```bash
$ csv2json -n file.csv
```

If the -n flag is not set and the CSV file contains a header row, the tool uses the header row to create JSON objects. The keys in the JSON objects correspond to the fields in the header row, and the values correspond to the data in each record.

```bash
$ csv2json file.csv
```

## Error Handling

csv2json handles different types of errors that can occur during CSV parsing.

If the tool encounters a record with a different number of fields than expected (as determined by the first record), it will print an error message and continue with the next record. It also handles and reports other types of errors while reading the CSV data.

## JSON Output

The JSON output is formatted with an indent of two spaces for better readability.

## Note

Be aware that csv2json will panic and exit immediately if it encounters an error while encoding to JSON. In such a case, check your CSV data for any irregularities.

## License

MIT License
