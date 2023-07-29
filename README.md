# JSON to Hosts Parser

![Banner](https://example.com/banner.png)

## Description

This tool is a Go (Golang) application that parses JSON data and extracts 
the host fields, writing unique domains to an output file.

## Installation

- Ensure you have Go language installed. [Go download 
page](https://golang.org/dl/)
- Clone this repo or download and extract the ZIP file.

## Usage

The application is run from the command line. Use the following parameters 
to specify the JSON data file and output file:

```bash
go run main.go -file <JSON_FILE> -output <OUTPUT_FILE>

