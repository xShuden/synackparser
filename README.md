# JSON to Hosts Parser

## Description

This tool is a Go (Golang) application that parses JSON data and extracts the host fields, writing unique domains to an output file.

## Installation

- Ensure you have Go language installed. [Go download page](https://golang.org/dl/)
- Clone this repo or download and extract the ZIP file.

## Usage

The application is run from the command line. Use the following parameters to specify the JSON data file and output file:

```bash
go run main.go --file <JSON_FILE> --output <OUTPUT_FILE>
```

Parameters:

- `--file` (required): Path to the JSON data file.
- `--output` (optional): Path to the output file. If not provided, a file named `output.txt` will be created.

You can also use the `--help` parameter to display the help page:

```bash
go run main.go --help
```

## Example Usage

```bash
go run main.go --file data.json --output domains.txt
```

## Requirements

- Go language installed
- Install the `github.com/fatih/color` package:
```bash
go get github.com/fatih/color
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contribution

If you wish to contribute to this project, please submit a pull request or open an issue. Contributions are welcome!

---

Copyright © 2023 [xShuden]
