package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
)

type IncludeEntry struct {
	Enabled  bool   `json:"enabled"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	File     string `json:"file"`
}

type TargetScope struct {
	Include []IncludeEntry `json:"include"`
	Exclude []interface{}  `json:"exclude"`
}

type Target struct {
	Scope TargetScope `json:"scope"`
}

type JSONData struct {
	Target Target `json:"target"`
}

func parseHostsFromJSON(jsonData []byte) ([]string, error) {
	var data JSONData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	hosts := make([]string, len(data.Target.Scope.Include))
	for i, entry := range data.Target.Scope.Include {
		host := strings.TrimSuffix(entry.Host, "$")
		host = strings.TrimPrefix(host, "^")
		host = strings.ReplaceAll(host, "\\", "")
		host = strings.ReplaceAll(host, ".+.", "")
		hosts[i] = host
	}

	return hosts, nil
}

func uniqueStrings(input []string) []string {
	keys := make(map[string]bool)
	var unique []string
	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			unique = append(unique, entry)
		}
	}
	return unique
}

func writeToTxtFile(filename string, data []string) error {
	content := strings.Join(data, "\n") + "\n"
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	return err
}

func printHelp() {
	fmt.Println("Usage: go run main.go --file <JSON_FILE> --output <OUTPUT_FILE>")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func main() {
	color.Yellow(" ________  ________  ________  ________  _______   ________")
	color.Yellow(" |\\   __  \\|\\   __  \\|\\   __  \\|\\   ____\\|\\  ___ \\ |\\   __  \\")
	color.Yellow(" \\ \\  \\|\\  \\ \\  \\|\\  \\ \\  \\|\\  \\ \\  \\___|\\ \\   __/|\\ \\  \\|\\  \\")
	color.Yellow("  \\ \\   __  \\ \\   _  _\\ \\  \\    \\ \\  \\  __\\ \\  \\_|/_\\ \\   _  _\\")
	color.Yellow("   \\ \\  \\ \\  \\ \\  \\\\  \\\\ \\  \\____\\ \\  \\|\\  \\ \\  \\_|\\ \\ \\  \\\\  \\|")
	color.Yellow("    \\ \\__\\ \\__\\ \\__\\\\ _\\\\ \\_______\\ \\_______\\ \\_______\\ \\__\\\\ _\\")
	color.Yellow("     \\|__|\\|__|\\|__|\\|__|\\|_______|\\|_______|\\|_______|\\|__|\\|__|")
	color.Yellow("")

	jsonFile := flag.String("file", "", "JSON data file")
	outputFile := flag.String("output", "", "Output file")
	showHelp := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *showHelp {
		printHelp()
		return
	}

	if *jsonFile == "" {
		fmt.Println("Error: JSON file path is missing. Please provide the JSON data file.")
		return
	}

	jsonData, err := ioutil.ReadFile(*jsonFile)
	if err != nil {
		fmt.Printf("Error: Unable to read JSON file '%s': %v\n", *jsonFile, err)
		return
	}

	hosts, err := parseHostsFromJSON(jsonData)
	if err != nil {
		fmt.Println("Error: Invalid JSON format. Please make sure you provide a valid JSON data.")
		return
	}

	uniqueHosts := uniqueStrings(hosts)

	if *outputFile == "" {
		*outputFile = "output.txt"
	}

	if err := writeToTxtFile(*outputFile, uniqueHosts); err != nil {
		fmt.Printf("Error: Failed to write to %s: %v\n", *outputFile, err)
		return
	}

	fmt.Printf("%d unique domains have been written to %s.\n", len(uniqueHosts), *outputFile)
}
