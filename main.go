package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type LogColor struct {
	Keyword string // The keyword to search for (in uppercase)
	Color   string // The ANSI 256-color escape code
}

func parseMappings(mappingStr string) ([]LogColor, error) {
	var logColors []LogColor
	var errs []string
	pairs := strings.Split(mappingStr, ",")
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}
		parts := strings.Split(pair, ":")
		if len(parts) != 2 {
			errs = append(errs, fmt.Sprintf("invalid mapping: %s", pair))
			continue
		}
		keyword := strings.ToUpper(strings.TrimSpace(parts[0]))
		colorCode := strings.TrimSpace(parts[1])
		// If the escape sequence isn't present, build it.
		if !strings.Contains(colorCode, "38;") {
			colorCode = "\033[38;5;" + colorCode + "m"
		}
		logColors = append(logColors, LogColor{Keyword: keyword, Color: colorCode})
	}
	if len(errs) > 0 {
		return logColors, fmt.Errorf(strings.Join(errs, "; "))
	}
	return logColors, nil
}

func main() {
	// Custom usage message for better UX.
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Reads from stdin and colors log lines based on keyword mappings.")
		fmt.Fprintln(os.Stderr, "\nOptions:")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n  %s -mappings=\"ERROR:196,WARNING:226,INFO:33\"\n", os.Args[0])
	}

	mappingArg := flag.String("mappings", "", "Comma-separated list of keyword:color mappings (e.g. ERROR:196,WARNING:226,INFO:33)")
	flag.Parse()

	var logColors []LogColor
	if *mappingArg != "" {
		var err error
		logColors, err = parseMappings(*mappingArg)
		if err != nil {
			// Warn the user but continue using the mappings that did parse.
			fmt.Fprintln(os.Stderr, "Warning:", err)
		}
		if len(logColors) == 0 {
			fmt.Fprintln(os.Stderr, "No valid mappings provided; exiting.")
			os.Exit(1)
		}
	} else {
		logColors = []LogColor{
			{Keyword: "ERROR", Color: "\033[38;5;196m"},   // Bright red
			{Keyword: "WARNING", Color: "\033[38;5;226m"}, // Bright yellow
			{Keyword: "INFO", Color: "\033[38;5;33m"},     // Blue
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		lineUpper := strings.ToUpper(line)
		appliedColor := ""
		for _, lc := range logColors {
			if strings.Contains(lineUpper, lc.Keyword) {
				appliedColor = lc.Color
				break
			}
		}
		if appliedColor != "" {
			fmt.Fprintf(writer, "%s%s\033[0m\n", appliedColor, line)
		} else {
			fmt.Fprintln(writer, line)
		}
		writer.Flush() // Ensure immediate output for each line.

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}
